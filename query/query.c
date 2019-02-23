#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct entry {
	int len;
	char *str;
	int weight;
	struct entry *next;
};

struct entry *read_input(void) {
	// NOTE: The format is %d %s %d
	struct entry *tmp;
	int tlen, tweight;
	char *tbuf;
	scanf("%d", &tlen);
	tbuf = malloc(tlen + 1);
	scanf("%s", tbuf);
	scanf("%d", &tweight);
	tmp = malloc(sizeof(struct entry));
	if (!tmp) return NULL;
	tmp->len = tlen;
	tmp->str = tbuf;
	tmp->weight = tweight;
	tmp->next = NULL;

	return tmp;
}

struct entry *slurp_input(int len) {
	struct entry *head, *it = NULL;
	head = read_input();  // set & remember head;
	len--;		      // adjust while loop
	it = head;
	while (len--) {
		it->next = read_input();
		it = it->next;
	}
	return head;
}

int main(int argc, char *argv[]) {
	int tclen;
	scanf("%d", &tclen);
	struct entry *head = slurp_input(tclen);
	// read query
	int qlen;
	scanf("%d", &qlen);
	char *query = malloc(qlen + 1);
	scanf("%s", query);

	char *ans = NULL;
	int aweight = -1;
	struct entry *it = head;
	while (it) {
		// if (strstr(it->str, query) && aweight < it->weight &&
		// strncmp(it->str, query, qlen) == 0) {
		if (strncmp(it->str, query, qlen) == 0 &&
		    aweight < it->weight) {
			ans = it->str;
			aweight = it->weight;
		}
		it = it->next;
	}

	printf("%s\n", ans);
}
