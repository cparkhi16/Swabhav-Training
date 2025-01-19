"use strict";
//const {Matcher} = require("./matcher");
const {filterEngine}= require("./filterEngine")
const {Subscription}=require("./subscriptionClasses")
const {registerSubscription} = require("./init");
var XMLHttpRequest = require('xhr2');
const {elemHide} = require("./elemHide");

function httpGetAsync(theUrl, callback)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() { 
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            callback(xmlHttp.responseText);
    }
    xmlHttp.open("GET", theUrl, true); // true for asynchronous 
    xmlHttp.send(null);
}
/*!
 * Parts of original code from ipv6.js <https://github.com/beaugunderson/javascript-ipv6>
 * Copyright 2011 Beau Gunderson
 * Available under MIT license <http://mths.be/mit>
 */

/**
 * Extracts host name from a URL.
 */
 function extractHostFromURL(/**String*/ url)
 {
   if (url && extractHostFromURL._lastURL == url)
     return extractHostFromURL._lastDomain;
 
   var host = "";
   try
   {
     host = new URI(url).host;
   }
   catch (e)
   {
     // Keep the empty string for invalid URIs.
   }
 
   extractHostFromURL._lastURL = url;
   extractHostFromURL._lastDomain = host;
   return host;
 }
 
 /**
  * Parses URLs and provides an interface similar to nsIURI in Gecko, see
  * https://developer.mozilla.org/en-US/docs/XPCOM_Interface_Reference/nsIURI.
  * TODO: Make sure the parsing actually works the same as nsStandardURL.
  * @constructor
  */
 function URI(/**String*/ spec)
 {
   this.spec = spec;
   this._schemeEnd = spec.indexOf(":");
   if (this._schemeEnd < 0)
     throw new Error("Invalid URI scheme");
 
   if (spec.substr(this._schemeEnd + 1, 2) != "//")
     this._hostPortStart = this._schemeEnd + 1;
   else
     this._hostPortStart = this._schemeEnd + 3;
 
   if (this._hostPortStart == spec.length)
     throw new Error("Empty URI host");
 
   this._hostPortEnd = spec.indexOf("/", this._hostPortStart);
   if (this._hostPortEnd < 0)
   {
     var queryIndex = spec.indexOf("?", this._hostPortStart);
     var fragmentIndex = spec.indexOf("#", this._hostPortStart);
     if (queryIndex >= 0 && fragmentIndex >= 0)
       this._hostPortEnd = Math.min(queryIndex, fragmentIndex);
     else if (queryIndex >= 0)
       this._hostPortEnd = queryIndex;
     else if (fragmentIndex >= 0)
       this._hostPortEnd = fragmentIndex;
     else
       this._hostPortEnd = spec.length;
   }
 
   var authEnd = spec.indexOf("@", this._hostPortStart);
   if (authEnd >= 0 && authEnd < this._hostPortEnd)
     this._hostPortStart = authEnd + 1;
 
   this._portStart = -1;
   this._hostEnd = spec.indexOf("]", this._hostPortStart + 1);
   if (spec[this._hostPortStart] == "[" && this._hostEnd >= 0 && this._hostEnd < this._hostPortEnd)
   {
     // The host is an IPv6 literal
     this._hostStart = this._hostPortStart + 1;
     if (spec[this._hostEnd + 1] == ":")
       this._portStart = this._hostEnd + 2;
   }
   else
   {
     this._hostStart = this._hostPortStart;
     this._hostEnd = spec.indexOf(":", this._hostStart);
     if (this._hostEnd >= 0 && this._hostEnd < this._hostPortEnd)
       this._portStart = this._hostEnd + 1;
     else
       this._hostEnd = this._hostPortEnd;
   }
 }
 URI.prototype =
 {
   spec: null,
   get scheme()
   {
     return this.spec.substring(0, this._schemeEnd).toLowerCase();
   },
   get host()
   {
     return this.spec.substring(this._hostStart, this._hostEnd);
   },
   get asciiHost()
   {
     var host = this.host;
     if (/^[\x00-\x7F]+$/.test(host))
       return host;
     else
       return punycode.toASCII(host);
   },
   get hostPort()
   {
     return this.spec.substring(this._hostPortStart, this._hostPortEnd);
   },
   get port()
   {
     if (this._portStart < 0)
       return -1;
     else
       return parseInt(this.spec.substring(this._portStart, this._hostPortEnd), 10);
   },
   get path()
   {
     return this.spec.substring(this._hostPortEnd);
   },
   get prePath()
   {
     return this.spec.substring(0, this._hostPortEnd);
   }
 };
 

// var responsefromHTTP = httpGetAsync("https://easylist-downloads.adblockplus.org/easylist.txt",(data)=>{
//     //console.log(" response from HTTP REQ  ",data)
//     let info={};
//     info.exists = true;
//     info.content = data;
//     let s= Subscription.fromURL("https://easylist-downloads.adblockplus.org/easylist.txt")
//     registerSubscription(s,info)
//     console.log("------------------loaded subs -----------",s)
//     //Matcher.hello("http://example.com/ad.png",4,"example.com","",false)

//      filterEngine.checkUrl("https://ad.doubleclick.net/ddm/trackclk/N2724.edmunds/B27451241.331349279;dc_trk_aid=524165283;dc_trk_cid=147563836;dc_lat=;dc_rdid=;tag_for_child_directed_treatment=;tfua=;ltd=",3,"ndtv.com","",false)
//     //filterEngine.checkUrl("http://materialmoon.com/abc/bge$third-party",3,"ndtv.com","",false)

// })

 let info={};
info.exists = true;
info.content = "[Adblock Plus 2.0]\n&ad_type=\n/ad?type=\n/ads?type=\n?adv_type=\n?type=ad&\n||tpc.googlesyndication.com^\n/pagead/1p-user-list/\n/ad/banner_^\nexample.com|"; //@@||ams.amazon.co.uk^$domain=ams.amazon.co.uk
//info.content = "[Adblock Plus 2.0]\n! Checksum: b1Gxz0wemare+xdHshRFCg\n! Last modified: 29 Aug 2022 11:21 UTC\n! Version: 202208291121\n||facebook.com/network_ads_common\n||facebook.com^*/instream/vast.xml?\n@@*.png\n||ad.doubleclick.net^"
//info.content ="[Adblock Plus 2.0]\n! Checksum: x8pt1NE/HsZVfdzRU5J6KA"
//info.content = `[Adblock Plus 2.0]\nexample.com##a[href^="http://ad.doubleclick.net/"]\nexample.com##a[href^="http://adf.ly/?id="]\n##div[aria-label="Ads"]\n##div\n##div[class^="AdCard_"]\n##.ad-300\nwsj.com##div[class*="WSJTheme--adWrapper"]\n##[id^="google_ads_iframe"]\n###top-banner-ad-browser\n##.img_ad`;
let s= Subscription.fromURL("https://easylist-downloads.adblockplus.org/easyprivacy.txt")
registerSubscription(s,info)
//console.log("------------------loaded subs -----------",s)
//filterEngine.checkUrl("https://googleadservices.com/page-ads-abc/aclk?sa=L&ai=CuibXstbsYoGVAo6amsMPsqajiAPX8djPa-aWltmREOH5pL79LBABIMzT5x5g5YKAgLwOoAGGyr-XAcgBCakC1N6wibIlSz7gAgCoAwHIAwqqBJ8CT9ARip9gbsPotvXnaozeGSuPJ8iK3zAoayOTgcnTMJVJpFveYz8wd6TPMn3ysa6IqEbv7qL2Mcs8adRYS3gv5TycUeT11AxQ9HpjNwxQAoy1EBJXnw-wIM7um8CWbiRmSVJxNtTygtSfVugoYCiV-TS6M3m9pnUkzrnS3AC8DF7QzX0CYMvB-1EjADnEZWdzCXT2DEzPnpuWqvnFjwL4BE_n2bqen67HaK--ORJNblZY0pVQzSfaNI3Zn0BeLw0lmHvRtplnC8Vz5guC_icgxDX68Dj1IWjiNVrYmHz78E1MhX-Y_VQjG7Njyjxwl_Jx0mlth8sjaG6Z30XXI7YFc06OI14zFdp3_lMQlOXMTwI8w3KLFX2oX2zJp3K_DB7ABJq2hbX1A-AEAZAGAaAGLoAH4rXA6AKIBwGQBwKoB47OG6gHk9gbqAfulrECqAf-nrECqAeko7ECqAfVyRuoB6a-G6gHmgaoB_PRG6gHltgbqAeqm7ECqAffn7EC2AcA0ggSCAAQAhhNMgEAOgef0ICAgIAEsQl8VKn7BxYKIoAKA5gLAcgLAYAMAbgMAbgT5APYEw2IFAHQFQGYFgH4FgGAFwE&ae=1&num=1&cid=CAMSeQClSFh3wPZUMctd-aLUCUgd3tuLUyw-bJPaD1tIGPLfpPTyYL8QD3Rse5CqpqtzMeSolTCJRdjPeOQ9_WcLG4wIfmOPbZI1WSfrlxE_c6xcJVtL-lmzB32wi9yXTIYMi3j4yL4lQ3j-CWaK4kD1CKDY-CItPcaY0Oc&sig=AOD64_20K73kl7p2C2M35FzDHvLFiy9MTQ&client=ca-pub-8609501543919728&nb=9&adurl=https://www.fireboltt.com%3Futm_source%3Dgoogleads%26utm_medium%3Ddisplay%26utm_campaign%3DGoogle_SW_Display_All_Conv_26072022%26gclid%3DCj0KCQjw_7KXBhCoARIsAPdPTfibqScNW5FcTH8vqFe_tzC-WXUH_6dtXqbU2vlgF4RxoEZVmXoXs90aAqWxEALw_wcB",3,"ndtv.com","",false)
//filterEngine.checkUrl("https://usmetric.rediff.com/www.rediff.com/dynimpression?rkey=980841&position=1_1-2_1&label=edit_topstories&c_type=edit_news&news=https%3a%2f%2fnews.rediff.com%2fcommentary%2f2022%2faug%2f24%2fliveupdates.htm",4,"","",false)
//filterEngine.checkUrl("https://abc.com/xyz/ad/banner_png",4,"","",false)
filterEngine.checkUrl("https://example.com",4,"","",false)
//filterEngine.checkUrl("https://ad.doubleclick.net/ddm/trackclk/N2724.edmunds/B27451241.331349279;dc_trk_aid=524165283;dc_trk_cid=147563836;dc_lat=;dc_rdid=;tag_for_child_directed_treatment=;tfua=;ltd=",4,"","",false)
// let url = "https://wsj.com"
// let host = url.indexOf(':') != -1 ? extractHostFromURL(url) : url;
// console.log(" host --== ",host)
// let a =  elemHide.getStyleSheet(host, false).code;
// console.log("code got ===>> ", a)

// initialize filter engine  -> initialize filterListener -> preload the subscription's _filterText array -> create a blocking filter from each of this filter in array and add it to the map in matcher -> use match for checking url and return matching filter if any

//updateSubscriptionFilters --=> starts entry for _filterText array and emits a event for creating a filter of each entry and deploying it


//@@ filtertext starting with @@ will be returned as a allowlisting filter










