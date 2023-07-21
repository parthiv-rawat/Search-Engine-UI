<script>
	import { onMount } from "svelte";
	import HomePage from "./HomePage.svelte";
	import Results from "./Results.svelte";

	let resultPage = false;
	let searchResults = [];
	let field = "";
	let signal = false;
	let searchQuery = "";
	let isProfile = false;
	let isHomePage = true;

	function showProfile() {
		isProfile = !isProfile;
	}

	function toggleIsHomePage() {
		if (isHomePage) {
			resultPage = false;
		}
	}

	async function cacheResults(query, results) {
		if (!("indexedDB" in window)) {
			console.error("IndexedDB is not supported in this browser.");
			return;
		}

		try {
			const openRequest = indexedDB.open("searchResultsStore", 1);

			openRequest.onupgradeneeded = (event) => {
				const dbInstance = event.target.result;
				if (!dbInstance.objectStoreNames.contains("cache")) {
					const store = dbInstance.createObjectStore("cache", {
						keyPath: "query",
					});
					store.createIndex("queryIndex", "query", {
						unique: true,
					});
				}
			};

			const db = await new Promise((resolve, reject) => {
				openRequest.onsuccess = (event) => {
					resolve(event.target.result);
				};

				openRequest.onerror = (event) => {
					reject(event.target.error);
				};
			});

			const tx = db.transaction("cache", "readwrite");
			const store = tx.objectStore("cache");
			store.add({ query, results});

			await new Promise((resolve) => {
				tx.oncomplete = resolve;
			});
			db.close();
		} catch (error) {
			console.error("Error caching data:", error);
		}
	}

	async function checkCache(query) {
		if (!("indexedDB" in window)) {
			console.error("IndexedDB is not supported in this browser.");
			return null;
		}

		try {
			const openRequest = indexedDB.open("searchResultsStore", 1);

			// console.log(openRequest)
			openRequest.onupgradeneeded = (event) => {
				const dbInstance = event.target.result;

				// Check if the object store already exists; if not, create it
				if (!dbInstance.objectStoreNames.contains("cache")) {
					const objectStore = dbInstance.createObjectStore("cache", {
						keyPath: "query",
					});
				}
			};
			// console.log(openRequest)
			const db = await new Promise((resolve, reject) => {
				openRequest.onsuccess = (event) => {
					resolve(event.target.result);
				};

				openRequest.onerror = (event) => {
					reject(event.target.error);
				};
			});

			const tx = db.transaction("cache", "readonly");
			const store = tx.objectStore("cache");

			// Check if the object store exists
			if (!db.objectStoreNames.contains("cache")) {
				console.error("Object store not found in the cache database.");
				return null;
			}

			const cachedData = await new Promise((resolve) => {
				const request = store.get(query);
				request.onsuccess = (event) => {
					resolve(event.target.result);
				};
				request.onerror = (event) => {
					console.error(
						"Error reading from cache:",
						event.target.error
					);
					resolve(null);
				};
			});

			db.close();

			return cachedData ? cachedData.results : null;
		} catch (error) {
			console.error("Error checking cache:", error);
			return null;
		}
	}

	async function handleSubmit(event) {
		event.preventDefault();
		// console.log(event)
		searchQuery = event.target.searchInput.value;
		// console.log(searchQuery);
		field = "";
		signal = true;

		// Check if data is cached
		const cachedData = await checkCache(searchQuery);
		console.log(cachedData)
		if (cachedData != null) {
			searchResults = cachedData;
			//   resultPage = true;
		} else {
			try {
				const resp = await fetch(
					`http://localhost:8081/search?query=${searchQuery}`
				);
				searchResults = await resp.json();
				// resultPage = true;
				// Cache the results
				cacheResults(searchQuery, searchResults);
			} catch (error) {
				console.error("Error fetching data:", error);
				searchResults = [];
			}
		}
		resultPage = true;
	}
</script>

{#if !resultPage && isHomePage}
	<HomePage {handleSubmit} on:click={showProfile} {isProfile} />
{:else}
	<Results
		{handleSubmit}
		{searchQuery}
		{searchResults}
		on:click={showProfile}
		{isProfile}
		on:click={toggleIsHomePage}
	/>
{/if}
