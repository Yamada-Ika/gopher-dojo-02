NAME	:= download
SRCS	:= main.go

all: $(NAME)

$(NAME): $(SRCS)
	go build -o $(NAME) $(SRCS)

clean:
	rm -rf $(NAME)

fclean: clean

re: fclean all

run:
	./download https://c.tenor.com/XCI9evLXjHcAAAAd/%E3%83%89%E3%83%A9%E3%81%88%E3%82%82%E3%82%93-doraemon.gif

rm:
	rm -rf *.gif

.PHONY: all clean fclean re test