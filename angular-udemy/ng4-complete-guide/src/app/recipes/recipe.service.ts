import { EventEmitter, Injectable } from "@angular/core";
import { Ingredient } from "../shared/ingredient.model";
import { ShoppingListService } from "../shopping-list/shopping-list.service";
import { Recipe } from "./recipe.model";
@Injectable()
export class RecipeService{
   constructor(private slService: ShoppingListService){}
    recipeSelected = new EventEmitter<Recipe>();
   private recipes:Recipe [] = [
        new Recipe('a test recipe','Test','https://cdn.pixabay.com/photo/2015/09/16/20/10/dough-943245_960_720.jpg',[new Ingredient('abc',1),new Ingredient('FF',4)]),
        new Recipe('another recipe','Test','https://cdn.pixabay.com/photo/2015/09/16/20/10/dough-943245_960_720.jpg',
        [new Ingredient('B',3),new Ingredient('C',5)])
      ];

    getRecipes(){
      //console.log("-=- recipes -= ",this.recipes)
        return this.recipes.slice();
    }
    addIngredientToShoppingList(ingredients: Ingredient[]){
      this.slService.addIngredients(ingredients)
    }
    getRecipe(index:number){
      return this.recipes[index]
    }
}