<html>

<head>
    <link rel="stylesheet" href="/assets/recherche.css" type="text/css">
    </link>
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Josefin+Slab">
    <title>Recherche de ville </title>
</head>

<body>
    <div class="centre-boite">
        <div class="boite-recherche">
            <form action="/recherche/ville" method="post">
                <fieldset>
                    <legend class="titre-recherche">Recherche de ville</legend>
                    <div class="form-recherche">
                        <label>Nom du maire <input type="text" name="owner"></label>

                        <label>Nom de la ville <input type="text" name="name"></label>

                        <input type="submit" value="Recherche">
                    </div>
                </fieldset>
            </form>
        </div>
    </div>
</body>

</html>