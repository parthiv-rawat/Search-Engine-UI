<script>
    import Pagination from "./Pagination.svelte";
    import { onMount, afterUpdate } from "svelte";
    import Profile from "./Profile.svelte";

    export let handleSubmit;
    export let searchQuery;
    let field = searchQuery;
    export let searchResults;
    export let isProfile;
    //  ----------------------------
    const itemsPerPage = 5;
    let currentPage = 1;
    let totalItems;

    if (searchResults !== "") {
        totalItems = searchResults;
    } else {
        totalItems = [1];
    }

    function onChangePage(newPage) {
        currentPage = newPage;
        updateDisplayedItems();
    }

    let displayedItems = [];

    onMount(() => {
        updateDisplayedItems();
    });

    afterUpdate(() => {
        totalItems = searchResults;
        updateDisplayedItems();
    });

    function updateDisplayedItems() {
        if (totalItems) {
            const startIndex = (currentPage - 1) * itemsPerPage;
            const endIndex = startIndex + itemsPerPage;
            displayedItems = totalItems.slice(startIndex, endIndex);
        }
    }
</script>

{#if !isProfile}
    <header>
        <a href="/"><img src="images/image.png" alt="Logo" class="logo" /></a>
        <form on:submit|preventDefault={handleSubmit}>
            <input
                type="text"
                id="search-input"
                name="searchInput"
                placeholder="Enter your search query..."
                bind:value={field}
            />
            <button type="submit"><img src="images/search.png" alt="" /></button
            >
        </form>
        <button type="toggle" on:click>Profile</button>
    </header>

    <main>
        <h3>Showing Results for {searchQuery}</h3>
        <div class="result">
            {#if typeof searchResults !== "string" && displayedItems.length > 0}
                {#each displayedItems as result}
                    <div class="blk">
                        <a href={result.website} target="_blank"
                            ><h3>{result.name}</h3></a
                        >
                        <p>{result.location}</p>
                        <p>{result.established}</p>
                        <p>{result.programs}</p>
                        {#each Object.entries(result.tuition) as [key, value]}
                            <p>{key}: {value}</p>
                        {/each}
                        {#each Object.entries(result.admission) as [key, value]}
                            <p>{key}: {value}</p>
                        {/each}
                    </div>
                {/each}
            {:else if typeof searchResults === "string" && searchResults !== ""}
                <h4>{searchResults}</h4>
            {:else if searchResults === ""}
                <h4>The search result does not exist.</h4>
            {:else}
                <p>No results found.</p>
            {/if}
        </div>
    </main>

    <Pagination
        {currentPage}
        totalPages={Math.ceil(totalItems.length / itemsPerPage)}
        {onChangePage}
    />

    <footer>
        <p>© 2023 Search Engine. All rights reserved.</p>
    </footer>
{:else}
    <Profile on:click />
{/if}

<style>
    img {
        display: inline;
        height: 50px;
        margin-left: 5px;
    }

    img:hover {
        scale: 1.2;
    }
    header {
        display: flex;
        justify-content: flex-start;
        align-items: center;
        background-color: #e9ab39;
        color: #fff;
        font-size: 10px;
        padding: 5px;
        text-align: center;
    }

    header button[type="toggle"] {
        position: absolute;
        right: 50px;
        padding: 3px;
        font-size: medium;
        background-color: #fff;
        border: 1px solid #c9c3c3;
        width: 55px;
        height: 55px;
        border-radius: 30%;
        display: block;
        cursor: pointer;
    }

    header button:hover {
        scale: 1.1;
    }
    /* h1 {
		display: inline;
		margin: 5px;
	} */

    header form {
        margin-left: 20px;
        display: flex;
        justify-content: center;
        align-items: center;
        /* margin-top: 30vh; */
    }

    input[type="text"] {
        margin: auto;
        padding: 10px;
        width: 70vw;
        border-radius: 30px;
        font-size: 16px;
        border: 0.5px solid #ccc;
    }

    input[type="text"]:hover {
        border: 1px solid #ccc;
        box-shadow: 3px 3px 4px 4px rgba(80, 78, 78, 0.2);
    }

    button[type="submit"] {
        /* margin: auto; */
        display: flex;
        margin-left: 8px;
        padding: 4px 10px;
        font-size: 16px;
        border-radius: 20px;
        background-color: #fff;
        color: #e9ab39;
        border: none;
        cursor: pointer;
    }

    button img {
        align-items: center;
        height: 25px;
        margin: 0;
        padding: 2px;
    }

    /* -------------------------------------------------------------------	 */
    main {
        min-height: 70.7vh;
        flex: 1;
    }
    main h3 {
        display: flex;
        margin-left: 14px;
        align-items: center;
        text-decoration: underline;
        justify-content: flex-start;
        text-decoration: none;
    }
    .result {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        justify-content: center;
    }

    /* .result h2 {
		text-decoration: underline;
		margin: 15px;
	} */

    .result div {
        margin: 10px;
    }

    h4 {
        margin-left: 15px;
    }

    .blk:hover {
        /* display: block; */
        padding: 3px;
        /* width: 40%; */
        scale: 1.05;
        /* border: 1px solid #ccc; */
        /* border-radius: 10px; */
        /* box-shadow: 2px 2px 2px 2px rgba(29, 28, 28, 0.4); */
        /* cursor: pointer; */
    }

    .blk a,
    .blk a:active {
        font-family: arial, sans-serif;
        font-weight: 400;
        text-decoration: none;
        color: #1a0dab;
    }

    .blk a:hover {
        text-decoration: underline;
    }

    .blk a > h3 {
        margin: 2px;
    }

    .blk > p {
        font-family: arial, sans-serif;
        margin: 2px;
    }

    footer {
        display: flex;
        flex-shrink: 0;
        bottom: 0;
        justify-content: center;
        align-items: center;
        margin-top: 5px;
        width: 98.5%;
        height: 10px;
        background-color: #e9ab39;
        color: #fff;
        padding: 10px;
        text-align: center;
    }
</style>
