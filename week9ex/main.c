#include<stdio.h>

void main()
{
	FILE* infile, * outfile;

	int i;
	double d;
	char str[32];

	infile = fopen("lsh1.txt", "r");
	outfile = fopen("lsh2.txt", "w");

	fscanf(infile, "%d%lf%s", &i, &d, str);

	printf("%d %lf %s\n", i, d, str);
	fprintf(outfile, "%d %lf %s\n", i, d, str);

	fclose(infile);
	fclose(outfile);
}