#include <stdio.h>

extern int A;
void SayHello(char name[]);

int main()
{
	printf("external variable A = %d\n", A);
	printf("external function call SayHello(Si-Hyung = ");
	SayHello("Si - Hyung");
	printf("external variable A = %d\n", A);
}