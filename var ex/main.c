#include <stdio.h>
#include <stdlib.h>

/*void main()
{
	int a, * ip, arr[3];

	a = 10;
	ip = &a;
	arr[0] = 10;
	a = arr[1];
	printf("a = %d\n", a);
}*/

void main()
{
	int* a;
	int* b;

	a = (int*)malloc(sizeof(int));
	b = (int*)malloc(sizeof(int));

	*a = 1;
	*b = 2;

	printf("*a = %d : *b = d\n", *a, *b);
	*a = *b;
	printf("*a = %d : *b = d\n", *a, *b);
}