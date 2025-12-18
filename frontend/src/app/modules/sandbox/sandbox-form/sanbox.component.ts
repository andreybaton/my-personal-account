import { ChangeDetectionStrategy, Component, Input, OnInit, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { BehaviorSubject, filter, map, Observable } from 'rxjs';
import { SandboxChildComponent } from "./sandboxChild/sanboxChild.component";
import { FormBuilder, FormGroup, UntypedFormBuilder, UntypedFormGroup } from '@angular/forms';


@Component({
  selector: 'sandbox',
  standalone: true,
  imports: [CommonModule, SandboxChildComponent],
  templateUrl: './sandbox.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush
//   styleUrl: './distlearning.css'
  
})
export class SandboxComponent  {
	
    @Output() mass: number[] = [];
    clicks = 0;
	eventForm: UntypedFormGroup | undefined;
	
	constructor(
		private _fb:UntypedFormBuilder
	){
		this.createForm();
		this.getDataForBackend()
		console.log(this.getDataForBackend());
	}
getDataForBackend() {
  const { options, event, address, ...rest } = this.eventForm?.value;
  return { ...rest, ...options, ...event, address };
}
	createForm() {
		this.eventForm =this._fb.group({
			state:[],
			options: this._fb.group({
				width:[],
				height:[]
			}),
			event: this._fb.group({
				date:[],
				place:[]
			}),
			address: this._fb.group({
				street:[],
				city:[]
			})
		});
	}

    onClick(){
        this.clicks++;
        //this.mass.push(this.clicks);
        //this.mass = [...this.mass, this.clicks]
    }
}
    // public subject = new BehaviorSubject<number>(10);
    // public obs$ = this.subject.asObservable();

    // public observable = new Observable<number>((subscriber)=>{
    //     subscriber.next(1);
    //     subscriber.next(2);
    //     subscriber.next(3);
    //     subscriber.next(4);
    //     subscriber.next(5);
    //     subscriber.complete();

    // });

    // ngOnInit(): void {
    //     this.observable.pipe(
    //         //map(num => num*num)
    //         filter(num=>num%2==0)
    //     ).subscribe({
    //         next: console.log});
    //     this.obs$.subscribe({
    //         next: console.log
    //     });

        
    // };
    

