import { Component, OnInit } from "@angular/core";
import { VerticalLayoutComponent } from "./layouts/vertical/vertical.component";

@Component({
    selector: 'layout',
    templateUrl: './layout.component.html',
    styleUrls: ['./layout.component.scss'],
    standalone   : true,
    imports: [VerticalLayoutComponent]
})

export class LayoutComponent implements OnInit{
    ngOnInit(): void {
        
    }
}