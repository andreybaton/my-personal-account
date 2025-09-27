import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, tap } from 'rxjs';

@Injectable({providedIn: 'root'})
export class DistlearningService
{

    private _data: BehaviorSubject<any> = new BehaviorSubject(null);

    // constructor(private _httpClient: HttpClient)
    // {
    // }

    // getData(): Observable<any>
    // {
    //     return this._httpClient.get('').pipe(
    //         tap((response: any) =>
    //         {
    //             this._data.next(response);
    //         }),
    //     );
    // }
}