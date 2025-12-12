import { AsyncPipe } from "@angular/common";
import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from "@angular/core";
import { Observable } from "rxjs";

@Component({
    selector: 'pending-button',
    templateUrl:'./pending-button.component.html',
    standalone: true,
    imports: [AsyncPipe],
    changeDetection: ChangeDetectionStrategy.Default
})
export class PendingButtonComponent {
    
    @Output() pendRequest = new EventEmitter<void>;
    status = 'pend';
    childClick(){
        console.log('childClick');
        this.pendRequest.emit()
    }
}