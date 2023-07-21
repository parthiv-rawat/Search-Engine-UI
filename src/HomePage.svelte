<script>
    import { onMount, onDestroy } from "svelte";
    import { fade, fly, scale } from "svelte/transition";
    import { quintIn, quintOut, elasticInOut } from "svelte/easing";
    import { tweened } from "svelte/motion";

    import InputAndSuggestions from "./InputAndSuggestions.svelte";
    import Profile from "./Profile.svelte";

    let field = "";
    let isVisible = true;
    export let isProfile;
    export let handleSubmit;
    
    // function toggleVisibility() {
    //     isVisible = !isVisible;
    // }
    // ----------------------------------------------------
</script>

{#if !isProfile}
    <header>
        <a href="/">
            <!-- {#if isVisible} -->
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <img src="images/image.png" alt="Logo" class="logo" />
            <!-- {/if} -->
            <h1>FlyHigh Edutech Search Engine</h1>
        </a>
        <button on:click>Profile</button>
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
{:else}
    <Profile on:click />
{/if}

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

    img:hover {
        scale: 1.1;
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

    header a {
        text-decoration: none;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    header a:visited {
        color: #fff;
    }
    header img {
        cursor: pointer;
    }

    header button {
        position: absolute;
        right: 50px;
        padding: 5px;
        font-size: medium;
        background-color: #fff;
        border: 1px solid #c9c3c3;
        width: 90px;
        height: 90px;
        border-radius: 30%;
        display: block;
        cursor: pointer;
    }

    header button:hover {
        scale: 1.15;
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
