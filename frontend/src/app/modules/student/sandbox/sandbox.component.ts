import { ChangeDetectionStrategy, Component } from "@angular/core";
import { PendingButtonComponent } from "./pending-button/pending-button.component";
import { delay, Observable, of } from "rxjs";
import { AsyncPipe } from "@angular/common";
import { SandboxService } from "./sandbox.service";

@Component({
    selector: 'sandbox-component',
    templateUrl:'./sandbox.component.html',
    standalone: true,
    changeDetection: ChangeDetectionStrategy.Default,
    imports:[PendingButtonComponent, AsyncPipe],
    providers:[SandboxService]
})
export class SandboxComponent {
    constructor(private sandServ: SandboxService){

    }

    pend(){
        console.log('pend')
        this.sandServ.getResponse().subscribe({
            next: (response)=>{
                console.log('res', response); 
            }
        })
    }
}