import { ChangeDetectionStrategy, Component, EventEmitter } from "@angular/core";
import { PendingButtonComponent } from "./pending-button/pending-button.component";
import { delay, Observable, of } from "rxjs";
import { AsyncPipe, NgForOf } from "@angular/common";
import { SandboxService } from "./sandbox.service";

@Component({
    selector: 'sandbox-component',
    templateUrl:'./sandbox.component.html',
    standalone: true,
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [PendingButtonComponent, AsyncPipe, NgForOf],
    providers:[SandboxService]
})
export class SandboxComponent {
    status = new EventEmitter<{id: number, status: string}>();
    //id: number = 1;
    constructor(private sandServ: SandboxService){

    }

    pend(ind: number){
        console.log('pend', ind)
        this.sandServ.getResponse().subscribe({
            next: (response)=>{
                console.log('res', response); 
                this.status.emit({id: ind, status: 'success'});
            }
        })
    }
}