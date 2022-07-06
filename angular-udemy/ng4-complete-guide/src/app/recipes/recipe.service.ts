import { EventEmitter } from "@angular/core";
import { Recipe } from "./recipe.model";

export class RecipeService{
    recipeSelected = new EventEmitter<Recipe>();
   private recipes:Recipe [] = [
        new Recipe('a test recipe','Test','https://cdn.pixabay.com/photo/2015/09/16/20/10/dough-943245_960_720.jpg'),
        new Recipe('another recipe','Test','https://cdn.pixabay.com/photo/2015/09/16/20/10/dough-943245_960_720.jpg')
    
      ];

    getRecipes(){
        return this.recipes.slice();
    }
}