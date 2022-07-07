import { Component, DoCheck, Input, OnInit } from '@angular/core';
import { Recipe } from '../recipe.model';
import { RecipeService } from '../recipe.service';

@Component({
  selector: 'app-recipe-detail',
  templateUrl: './recipe-detail.component.html',
  styleUrls: ['./recipe-detail.component.css']
})
export class RecipeDetailComponent implements OnInit,DoCheck {
  @Input() recipe: Recipe;
  constructor(private recipeService : RecipeService) { }
  ngDoCheck(): void {
  }
  ngOnInit(): void {
  }
  onAddToShoppingList(){
    this.recipeService.addIngredientToShoppingList(this.recipe.ingredients)
    console.log("--- ",this.recipe.ingredients)
  }

}
