#include <stdio.h>
#include <stdlib.h>

void main(int argc, char** argv)	//void main(int argc, char* argv[])	//void main(int argc, char argv[][])
{
	FILE* in;
	char name[32];
	int m_score, f_score;
	for (int i = 0; i < argc; i++) {
		printf("argv[%d] = %s\n", i, argv[i]);
	}
	//exit(1);
	in = fopen(argv[1], "r");
	fscanf(in, "%s%d%d", name, &m_score, &f_score);
	printf("%s %d %d", name, m_score, f_score);
	
	fclose(in);
}