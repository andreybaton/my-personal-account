import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { delay, Observable, of } from "rxjs";

@Injectable({providedIn:'root'})
export class SandboxService {
    constructor(_http: HttpClient){

    }

    getResponse(): Observable<string> {
        return of('string from serv').pipe(
            delay(2000)
        );
    }

}