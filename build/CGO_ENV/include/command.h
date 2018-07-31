#ifndef COMMAND_H
#define COMMAND_H

typedef struct
{
	char *command;
	void (*fun)(int argc, char *argv[], FILE *pfin, FILE *pfout);
}command_t;  

#define MAX_STRING 1024

const char* get_filename(const char* fullpathname, char *path);
unsigned long get_file_size(const char *filename, FILE *pfout);


typedef void (*out_callback_fun)(void *fg, char out[], int n);
int run_cmd_quiet(FILE *pfout, out_callback_fun fn, void *fg, const char * format, ...);

#define FAILED_OUT_HELPER(fmt, ...) fprintf(pfout, "failed["fmt"]\n", ##__VA_ARGS__)
#define FAILED_OUT(...)	FAILED_OUT_HELPER(__VA_ARGS__)
#define SUCESS_OUT()	fprintf(pfout, "sucess\n")
#define PROGRESS_OUT_HELPER(p, fmt, ...)	fprintf(pfout, "progress %d%%["fmt"]\n", p, ##__VA_ARGS__)
#define PROGRESS_OUT(p,...)	PROGRESS_OUT_HELPER(p, __VA_ARGS__)

int deal_command(char *cmdline, FILE *pfin, FILE *pfout);

#define BUILDIN_CMD(c,f)	static command_t f##_buildin_cmd \
	__attribute__((section("buildin_cmd"),__used__,aligned(sizeof(void*)))) = {c, f}

#endif //command.h
