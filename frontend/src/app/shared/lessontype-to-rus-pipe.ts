import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'lessontypeToRus'
})
export class LessontypeToRusPipe implements PipeTransform {

  transform(value: unknown, ...args: unknown[]): unknown {
	switch(value){
		case('lecture'):
			return 'Лекция';
		case('practice'):
			return 'Практика';
		case('lab'):
			return 'Лабораторная работа';
		default:
			return 'Тип не указан';
	}
    return null;
  }

}
