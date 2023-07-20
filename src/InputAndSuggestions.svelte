<script>
    import { afterUpdate } from "svelte";
    import { debounce } from "lodash";
    let showSuggestions = false;
    let suggestions = [];
    let searchSuggestions = [];
    let promiseSuggestions;
    let field = "";
    let inputElement;

    let debouncedHandleInput = null;
    afterUpdate(() => {
        debouncedHandleInput = debounce(handleInput, 300);
    });

    function handleInput(event) {
        const inputValue = event.target.value;
        field = inputValue;
        getAutocompleteSuggestions(inputValue);
    }

    function selectSuggestion(suggestion) {
        field = suggestion.title;
        showSuggestions = false;
        inputElement.focus();
    }

    async function getAutocompleteSuggestions(inputValue) {
        try {
            const resp = await fetch(
                `http://localhost:8081/search?query=${inputValue}`
            );
            searchSuggestions = await resp.json();
        } catch (error) {
            console.error("Error fetching data:", error);
            searchSuggestions = [];
        }

        suggestions = searchSuggestions;
        // console.log(Object.entries(suggestions[0]));
        showSuggestions =
            suggestions.length > 0 &&
            suggestions.length != 34;
        console.log(showSuggestions);
    }
</script>

<input
    type="text"
    id="search-input"
    name="searchInput"
    placeholder="Enter your search query..."
    on:input={() => debouncedHandleInput(event)}
    bind:value={field}
    bind:this={inputElement}
/>
<button type="submit" on:click><img src="images/search.png" alt="" /></button>
<div class="autocomplete">
    {#if showSuggestions}
        <ul>
            {#each suggestions as suggestion}
                <li on:click={() => selectSuggestion(suggestion)}>
                    {suggestion.title}
                </li>
            {/each}
        </ul>
    {/if}
</div>

<style>
    input[type="text"] {
        padding: 10px;
        width: 150vh;
        border-radius: 50px;
        font-size: 16px;
        border: 1px solid #ccc;
        box-shadow: 3px 3px 4px rgba(29, 28, 28, 0.1);
    }

    button[type="submit"] {
        margin-left: 20px;
        padding: 10px 25px;
        font-size: 16px;
        border-radius: 40px;
        background-color: #e9ab39;
        color: #fff;
        border: none;
        cursor: pointer;
    }

    button img {
        height: 30px;
        padding: 2px;
    }

    .autocomplete {
        position: absolute;
        width: 70%; /* Adjust the width as needed */
        top: 58vh;
        left: 25vh;
    }

    .autocomplete ul {
        position: absolute;
        top: 10%;
        left: 0;
        right: 0;
        background-color: #fff;
        border: 1px solid #ccc;
        border-radius: 24px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        max-height: 200px;
        overflow-y: auto;
        z-index: 10; /* Ensure the suggestions appear above other elements */
        padding: 0; /* Remove default padding for the list */
    }

    .autocomplete ul li {
        padding: 8px 12px;
        cursor: pointer;
        list-style: none; /* Remove the default list bullet */
        transition: background-color 0.2s; /* Add a transition effect */
    }

    .autocomplete ul li:not(:last-child) {
        border-bottom: 1px solid #ccc;
    }

    .autocomplete ul li:hover {
        background-color: #f0f0f0; /* Highlight the hovered suggestion */
    }

    .autocomplete ul li.active {
        background-color: #e9ab39; /* Highlight the selected suggestion */
        color: #fff; /* Set text color for better contrast */
    }

    /* Custom scrollbar for the suggestions list */
    .autocomplete ul::-webkit-scrollbar {
        width: 8px;
    }

    .autocomplete ul::-webkit-scrollbar-thumb {
        background-color: rgba(0, 0, 0, 0.2);
        border-radius: 4px;
    }
</style>
