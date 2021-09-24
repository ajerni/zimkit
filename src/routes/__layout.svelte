<!-- src/routes/__layout.svelte wird für alle pages angewandt -->
<!-- https://rodneylab.com/getting-started-with-sveltekit/ -->

<header>
    <nav>
            <a href="/.">Home</a>
            <a href="/contact">Contact</a>
            <a href="/temperatur">Temperatur</a>
    </nav>
</header>
<main>
<slot />
</main>
<footer>
    <p>
            A project by
            <a aria-label="Open the Rodney Lab site" href="https://www.andierni.ch/" target="_blank" rel="nofollow noopener noreferrer"
                    >ERNILABS</a
            >.
    </p>
</footer>