#include<stdio.h>
#include<string.h>

#define A 100;
#define PI 3.14;
#define C 'K';
#define STR "HANSUNG"

void main()
{
	int i;
	double pi;
	char c;
	char str[256];

	i = A;
	pi = PI;
	c = C;
	strcpy(str, STR);

	printf("i = %d\n", i);
	printf("d = %lf\n", pi);
	printf("c = %c\n", c);
	printf("c = %d\n", c);
	printf("str = %s\n", str);
}