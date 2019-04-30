// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_TEMPLATE_UTILS_H
#define MCSERVER_TEMPLATE_UTILS_H

#define VAR_LEN                 64

// 模板变量
typedef struct template_var template_var;
struct template_var {
    char *var_name;
    char *var_value;
    int v_len;
};

/**
 * 检查并获取
 * @param str
 * @param var_name
 * @return
 */
int check_and_get_var_name(const char *str, char *var_name);


// 模板解析
// @src_temple:模板
// @len:模板长度
// @vars:动态变量
// @output:输出解析后的字符串
void parse_template(const char *src_temple, int len, template_var vars[], int vars_len, char *output);

/**
 * 获取rbctrl.sh模板
 * @return
 */
char *get_rbctrl_sh_template();

/**
 * 获取interfaces模板
 * @return
 */
char *get_interfaces_template();


#endif //MCSERVER_TEMPLATE_UTILS_H
