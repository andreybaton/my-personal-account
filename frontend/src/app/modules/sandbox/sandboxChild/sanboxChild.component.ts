import { ChangeDetectionStrategy, Component, Input, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { BehaviorSubject, filter, map, Observable } from 'rxjs';


@Component({
  selector: 'sandbox-child',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './sandboxChild.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush
//   styleUrl: './distlearning.css'
  
})
export class SandboxChildComponent  {
    @Input() mass: number[] = [];
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
    

}