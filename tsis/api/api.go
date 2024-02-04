package api

type Movie struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Director    string  `json:"director"`
	Genre       string  `json:"genre"`
	ReleaseYear int     `json:"release_year"`
	Rating      float64 `json:"rating"`
}

var Movies = []Movie{
	{1, "Inception", "Christopher Nolan", "Sci-Fi", 2010, 8.8},
	{2, "The Shawshank Redemption", "Frank Darabont", "Drama", 1994, 9.3},
	{3, "The Godfather", "Francis Ford Coppola", "Crime", 1972, 9.2},
	{4, "The Dark Knight", "Christopher Nolan", "Action", 2008, 9.0},
	{5, "Schindler's List", "Steven Spielberg", "Biography", 1993, 8.9},
	{6, "Forrest Gump", "Robert Zemeckis", "Drama", 1994, 8.8},
	{7, "The Lord of the Rings: The Return of the King", "Peter Jackson", "Adventure", 2003, 8.9},
	{8, "Pulp Fiction", "Quentin Tarantino", "Crime", 1994, 8.9},
	{9, "Fight Club", "David Fincher", "Drama", 1999, 8.8},
	{10, "The Lord of the Rings: The Fellowship of the Ring", "Peter Jackson", "Adventure", 2001, 8.8},
	{11, "The Matrix", "Lana Wachowski, Lilly Wachowski", "Action", 1999, 8.7},
	{12, "The Silence of the Lambs", "Jonathan Demme", "Crime", 1991, 8.6},
	{13, "The Lord of the Rings: The Two Towers", "Peter Jackson", "Adventure", 2002, 8.7},
	{14, "Se7en", "David Fincher", "Crime", 1995, 8.6},
	{15, "Inglourious Basterds", "Quentin Tarantino", "Adventure", 2009, 8.3},
	{16, "Goodfellas", "Martin Scorsese", "Biography", 1990, 8.7},
	{17, "The Godfather: Part II", "Francis Ford Coppola", "Crime", 1974, 9.0},
	{18, "The Shining", "Stanley Kubrick", "Drama", 1980, 8.4},
	{19, "The Pianist", "Roman Polanski", "Biography", 2002, 8.5},
	{20, "Gladiator", "Ridley Scott", "Action", 2000, 8.5},
	{21, "The Departed", "Martin Scorsese", "Crime", 2006, 8.5},
	{22, "The Dark Knight Rises", "Christopher Nolan", "Action", 2012, 8.4},
	{23, "The Green Mile", "Frank Darabont", "Crime", 1999, 8.6},
	{24, "Django Unchained", "Quentin Tarantino", "Drama", 2012, 8.4},
	{25, "The Lion King", "Roger Allers, Rob Minkoff", "Animation", 1994, 8.5},
	{26, "The Godfather: Part III", "Francis Ford Coppola", "Crime", 1990, 7.6},
	{27, "Braveheart", "Mel Gibson", "Biography", 1995, 8.3},
	{28, "The Prestige", "Christopher Nolan", "Drama", 2006, 8.5},
	{29, "A Beautiful Mind", "Ron Howard", "Biography", 2001, 8.2},
	{30, "The Usual Suspects", "Bryan Singer", "Crime", 1995, 8.5},
	{31, "Interstellar", "Christopher Nolan", "Adventure", 2014, 8.6},
	{32, "The Intouchables", "Olivier Nakache, Eric Toledano", "Biography", 2011, 8.5},
	{33, "The Dark Knight Rises", "Christopher Nolan", "Action", 2012, 8.4},
	{34, "Pulp Fiction", "Quentin Tarantino", "Crime", 1994, 8.9},
	{35, "Forrest Gump", "Robert Zemeckis", "Drama", 1994, 8.8},
}
