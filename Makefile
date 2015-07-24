all: recipes
	sudo docker run -p 8000:8000 recipes
recipes: .
	sudo docker build -t recipes .
