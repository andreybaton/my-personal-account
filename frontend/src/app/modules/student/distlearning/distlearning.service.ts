import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, tap } from 'rxjs';
import { Lesson } from '../../../shared/lesson.interface';

@Injectable({providedIn: 'root'})
export class DistlearningService
{

    private _data: BehaviorSubject<any> = new BehaviorSubject(null);

    private apiUrl = 'https://localhost:7179/api/lessons'; // для HTTPS
    // или для HTTP: 'http://localhost:5281/api/lessons'

	constructor(private http: HttpClient) { }

 	// Получить все уроки
	getAllLessons(): Observable<Lesson[]> {
		//console.log(this.http.get<Lesson[]>(this.apiUrl));
		return this.http.get<Lesson[]>(this.apiUrl);
	}
}