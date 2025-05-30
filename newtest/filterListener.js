/*
 * This file is part of Adblock Plus <https://adblockplus.org/>,
 * Copyright (C) 2006-present eyeo GmbH
 *
 * Adblock Plus is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * Adblock Plus is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Adblock Plus.  If not, see <http://www.gnu.org/licenses/>.
 */

/** @module */

"use strict";

/**
 * @file Synchronization between filter storage and the filter engine.
 */

const {filterStorage} = require("./filterStorage");
const {filterNotifier} = require("./filterNotifier");
const {isActiveFilter, Filter} = require("./filterClasses");

/**
 * Checks whether filters from a given subscription should be deployed to the
 * filter engine.
 *
 * If the subscription is both valid and enabled, the function returns `true`;
 * otherwise, it returns `false`.
 *
 * @param {module:subscriptionClasses.Subscription} subscription
 *   The subscription.
 *
 * @returns {boolean} Whether filters from the subscription should be deployed
 *   to the filter engine.
 */
function shouldDeployFilters(subscription) {
  //console.log(" s valid disbaled -= ",subscription.valid, subscription.disabled)
  return subscription.valid && !subscription.disabled;
}

/**
 * Deploys a filter to the filter engine.
 *
 * The filter is deployed only if it belongs to at least one subscription that
 * is both valid and enabled.
 *
 * If the filter is a snippet filter, or a block filter with the `$header`
 * options, it is deployed only if it belongs to at
 * least one subscription that is valid, enabled, and of
 * {@link module:subscriptionClasses.Subscription#type type}
 * `circumvention` or a
 * {@link module:subscriptionClasses.SpecialSubscription special subscription}
 * that keeps user-defined filters.
 *
 * @param {module:filterEngine~FilterEngine} engine A reference to the
 *   filter engine.
 * @param {Filter} filter The filter.
 * @param {?Array.<module:subscriptionClasses.Subscription>} [subscriptions]
 *   A list of subscriptions to which the filter belongs. If omitted or `null`,
 *   the information is looked up from
 *   {@link module:filterStorage.filterStorage filter storage}.
 */
function deployFilter(engine, filter, subscriptions = null) {
  //console.log(" filter pattern , ",filter.pattern , filter.requiresPrivilegedSubscription)
  if (!isActiveFilter(filter))
    return;

  let deploy = false;
  let allowPrivileged = false;
  let requiresPrivilegedSubscription = filter.requiresPrivilegedSubscription;

  for (let subscription of subscriptions ||
                           filterStorage.subscriptions(filter.text)) {
    if (shouldDeployFilters(subscription) &&
        !filter.isDisabledForSubscription(subscription.url)) {
      deploy = true;
      if (!requiresPrivilegedSubscription)
        break;

      // Allow snippets to be executed only by the circumvention lists or the
      // user's own filters.
      if (subscription.privileged) {
        allowPrivileged = true;
        break;
      }
    }
  }

  if (!deploy)
    return;

  if (requiresPrivilegedSubscription && !allowPrivileged)
    return;
  //console.log(" adding filter -============== >>>>",filter.pattern,requiresPrivilegedSubscription,allowPrivileged)
  engine.add(filter);
}

/**
 * Undeploys a filter from the filter engine.
 *
 * The filter is undeployed only if it does not belong to at least one
 * subscription that is both valid and enabled.
 *
 * @param {module:filterEngine~FilterEngine} engine A reference to the
 *   filter engine.
 * @param {module:filterClasses.Filter} filter The filter.
 */
function undeployFilter(engine, filter) {
  if (!isActiveFilter(filter))
    return;

  let requiresPrivilegedSubscription = filter.requiresPrivilegedSubscription;

  for (let subscription of filterStorage.subscriptions(filter.text)) {
    if (shouldDeployFilters(subscription) &&
        !filter.isDisabledForSubscription(subscription.url)) {
      if (!requiresPrivilegedSubscription ||
          (requiresPrivilegedSubscription && subscription.privileged))
        return;
    }
  }

  engine.remove(filter);
}

/**
 * `{@link module:filterListener.filterListener filterListener}`
 * implementation.
 */
class FilterListener {
  /**
   * Initializes filter listener on startup, registers the necessary hooks.
   *
   * Initialization is asynchronous; once complete,
   * `{@link module:filterNotifier.filterNotifier filterNotifier}` emits the
   * `ready` event.
   *
   * @hideconstructor
   */
  constructor() {
    /**
     * A reference to the filter engine.
     * @type {?module:filterEngine~FilterEngine}
     * @private
     */
    this._engine = null;

    /**
     * Increases on filter changes, filters will be saved if it exceeds 1.
     * @type {number}
     * @private
     */
    this._isDirty = 0;
  }

  /**
   * Initializes filter listener.
   * @param {module:filterEngine~FilterEngine} engine A reference to the
   *   filter engine.
   * @returns {Promise} A promise that is fulfilled when the initialization is
   *   complete.
   * @package
   */
   initialize(engine) {
    //console.log("within initialize of filterListener ===")
    if (engine == null || typeof engine != "object")
      throw new Error("engine must be a non-null object.");

    if (this._engine != null)
      throw new Error("Filter listener already initialized.");

    this._engine = engine;

    filterStorage.loadFromDisk();
    
    //let promise = Promise.resolve();
    //console.log(" Initialize filters from each subscription asynchronously on startup by")
    // Initialize filters from each subscription asynchronously on startup by
    // setting up a chain of promises.
    for (let subscription of filterStorage.subscriptions()) {
      //console.log(" here in initialize of filterListener ",subscription)
      //console.log( " my  subs ",subscription)
     // subscription.valid = true
     // if (shouldDeployFilters(subscription)) {
        //await promise.then(() => {
          for (let text of subscription.filterText())
          //console.log(text)
            deployFilter(this._engine, Filter.fromText(text), [subscription]);
        //});
      //}
    }
    //console.log(" after for loop --------->>")
    filterNotifier.on("filter.added", this._onFilterAdded.bind(this));
    filterNotifier.on("filter.removed", this._onFilterRemoved.bind(this));
    // filterNotifier.on("filter.moved", this._onGenericChange.bind(this));

    filterNotifier.on("filterState.enabled",
                      this._onFilterStateEnabled.bind(this));
    filterNotifier.on("filterState.disabledSubscriptions",
                      this._onFilterStateEnabled.bind(this));
    filterNotifier.on("filterState.hitCount",
                      this._onFilterStateHitCount.bind(this));
    filterNotifier.on("filterState.lastHit",
                      this._onFilterStateLastHit.bind(this));

    filterNotifier.on("subscription.added",
                      this._onSubscriptionAdded.bind(this));
    filterNotifier.on("subscription.removed",
                      this._onSubscriptionRemoved.bind(this));
    filterNotifier.on("subscription.disabled",
                      this._onSubscriptionDisabled.bind(this));
    filterNotifier.on("subscription.updated",
                      this._onSubscriptionUpdated.bind(this));
    //filterNotifier.on("subscription.title", this._onGenericChange.bind(this));
    // filterNotifier.on("subscription.fixedTitle",
    //                   this._onGenericChange.bind(this));
    // filterNotifier.on("subscription.homepage",
    //                   this._onGenericChange.bind(this));
    // filterNotifier.on("subscription.downloadStatus",
    //                   this._onGenericChange.bind(this));
    // filterNotifier.on("subscription.lastCheck",
    //                   this._onGenericChange.bind(this));
    // filterNotifier.on("subscription.errors",
    //                   this._onGenericChange.bind(this));

    filterNotifier.on("load", this._onLoad.bind(this));
    filterNotifier.on("save", this._onSave.bind(this));

    // Indicate that all filters are ready for use.
    filterNotifier.emit("ready");
  }

  /**
   * Increases "dirty factor" of the filters and calls
   * filterStorage.saveToDisk() if it becomes 1 or more.
   *
   * Save is executed delayed to prevent multiple subsequent calls. If the
   * parameter is 0 it forces saving filters if any changes were recorded after
   * the previous save.
   *
   * @param {number} factor
   *
   * @private
   */
  // _setDirty(factor) {
  //   if (factor == 0 && this._isDirty > 0)
  //     this._isDirty = 1;
  //   else
  //     this._isDirty += factor;
  //   if (this._isDirty >= 1) {
  //     this._isDirty = 0;
  //     filterStorage.saveToDisk();
  //   }
  // }

  _onSubscriptionAdded(subscription) {
    //this._setDirty(1);
    //console.log("_onSubscriptionAdded >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
    if (shouldDeployFilters(subscription)) {
      for (let text of subscription.filterText()){
      //console.log("within if _onSubscriptionAdded")
        deployFilter(this._engine, Filter.fromText(text), [subscription]);
      }
    }
  }

  _onSubscriptionRemoved(subscription) {
    this._setDirty(1);

    if (shouldDeployFilters(subscription)) {
      for (let text of subscription.filterText())
        undeployFilter(this._engine, Filter.fromText(text));
    }
  }

  _onSubscriptionDisabled(subscription, newValue) {
    this._setDirty(1);
    //console.log("_onSubscriptionDisabled")
    if (filterStorage.hasSubscription(subscription)) {
      if (newValue == false) {
        for (let text of subscription.filterText())
          deployFilter(this._engine, Filter.fromText(text), [subscription]);
      }
      else {
        for (let text of subscription.filterText())
          undeployFilter(this._engine, Filter.fromText(text));
      }
    }
  }

  _onSubscriptionUpdated(subscription, textDelta) {
   // this._setDirty(1);
    console.log("_onSubscriptionUpdated  >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
    if (shouldDeployFilters(subscription) &&
        filterStorage.hasSubscription(subscription)) {
      for (let text of textDelta.removed)
        undeployFilter(this._engine, Filter.fromText(text));

      for (let text of textDelta.added)
      {
          console.log(" kk ",Filter.fromText(text))
        deployFilter(this._engine, Filter.fromText(text), [subscription]);
      }
    }
  }

  _onFilterAdded(filter) {
    this._setDirty(1);
    console.log("_onFilterAdded")
    deployFilter(this._engine, filter);
  }

  _onFilterRemoved(filter) {
    this._setDirty(1);
    undeployFilter(this._engine, filter);
  }

  _onFilterStateEnabled(text, newValue) {
    this._setDirty(1);

    // it's difficult to tell from newValue alone if we need to add or
    // remove the filter. Luckily, deployFilter and undeployFilter do
    // full checks that the filter should be added / removed.
    let filter = Filter.fromText(text);
    if (this._engine.has(filter))
      undeployFilter(this._engine, filter);
    else
    {
      deployFilter(this._engine, filter);
      console.log("_onFilterStateEnabled")
    }
  }

  _onFilterStateHitCount(text, newValue) {
    if (newValue == 0)
      this._setDirty(0);
    else
      this._setDirty(0.002);
  }

  _onFilterStateLastHit() {
    this._setDirty(0.002);
  }

  _onGenericChange() {
    this._setDirty(1);
  }

  _onLoad() {
    this._isDirty = 0;

    this._engine.clear();
    console.log("_onLoad")
    for (let subscription of filterStorage.subscriptions()) {
      if (shouldDeployFilters(subscription)) {
        for (let text of subscription.filterText())
          deployFilter(this._engine, Filter.fromText(text), [subscription]);
      }
    }
  }

  _onSave() {
    this._isDirty = 0;
  }
}

/**
 * Component synchronizing filter storage with the filter engine.
 * @type {module:filterListener~FilterListener}
 * @package
 */
exports.filterListener = new FilterListener();
