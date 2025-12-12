import { AsyncPipe } from "@angular/common";
import { ChangeDetectionStrategy, Component, EventEmitter, Input, OnInit, Output } from "@angular/core";
import { BehaviorSubject, Observable, timer } from "rxjs";

@Component({
    selector: 'pending-button',
    templateUrl:'./pending-button.component.html',
    styleUrl:'./pending-button.component.scss',
    standalone: true,
    imports: [AsyncPipe],
    changeDetection: ChangeDetectionStrategy.Default
})
export class PendingButtonComponent implements OnInit {
    @Input() butIndex: any;
    @Input() statusUpdate!: EventEmitter<{id: number, status: string}>;
    @Output() pendRequest = new EventEmitter<void>;
    statuses: 'default'|'pending'|'success'|'error' = 'default';
    status = new BehaviorSubject(this.statuses);
    status$ = this.status.asObservable();
    
    ngOnInit(): void {
        this.statusUpdate.subscribe((value)=>{this.handleStatus(value.id, value.status)})
    }

    childClick(ind: any){
        console.log('childClick');
        this.status.next('pending');
        this.pendRequest.emit(ind)
    }

    handleStatus(ind:number, status: string){
        if (this.butIndex == ind){
            switch(status){
                case('success'):
                this.status.next('success');
                setTimeout(()=>{this.status.next('default')},2000);

            }
        }
    }
}