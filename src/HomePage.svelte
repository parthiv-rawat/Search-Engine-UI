<script>
    
    import { fade, fly, scale } from "svelte/transition";
    import { quintIn, quintOut, elasticInOut } from "svelte/easing";
    import { tweened } from "svelte/motion";
    import InputAndSuggestions from "./InputAndSuggestions.svelte";

    let field = "";
    let isVisible = true;
    export let handleSubmit;

    // function toggleVisibility() {
    //     isVisible = !isVisible;
    // }
    // ----------------------------------------------------
    
</script>

<header>
    <!-- {#if isVisible} -->
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <img src="images/image.png" alt="Logo" class="logo" />
    <!-- {/if} -->
    <h1>FlyHigh Edutech Search Engine</h1>
</header>

<main>
    <!-- {#if !isVisible} -->
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- <img
        src="images/image.png"
        on:click={toggleVisibility}
            transition:scale
            alt="Logo"
            class="logo"
            /> -->
    <!-- {/if} -->
    <form on:submit={handleSubmit}>
       <InputAndSuggestions />
    </form> 
</main>

<footer>
    <p>© 2023 Search Engine. All rights reserved.</p>
</footer>

<style>
    :root {
        --margin: 10vh;
    }

    img {
        display: inline;
        height: 100px;
    }
    header {
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: #e9ab39;
        color: #fff;
        padding: 5px;
        text-align: center;
    }

    header img {
        cursor: pointer;
    }
    h1 {
        display: inline;
        margin: 5px;
    }

    main {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
        height: 64vh;
        flex: 1;
        /* margin-top: 5vh; */
        padding: 20px;
        text-align: center;
    }

    /* main > img {
        display: flex;
        cursor: pointer; 
        margin-bottom: var(--margin); 
        width: 180px;
        height: 180px;
        justify-content: center;
        align-items: center;
    } */

    
    footer {
        display: flex;
        flex-shrink: 0;
        bottom: 0;
        justify-content: center;

        align-items: center;
        width: 98.5%;
        height: 20px;
        background-color: #e9ab39;
        color: #fff;
        padding: 10px;
        text-align: center;
    }

    form {
        display: flex;
        justify-content: center;
        margin-top: calc(40vh - var(--margin));
    }

</style>
