<!-- Pagination.svelte -->
<script>
    export let currentPage;
    export let totalPages;
    export let onChangePage;
    // console.log(totalPages)
</script>

<div class="pagination">
    {#if currentPage > 1}
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <div class="page" on:click={() => onChangePage(1)}>First</div>
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <div class="page" on:click={() => onChangePage(currentPage - 1)}>
            Previous
        </div>
    {/if}

    {#if currentPage > 3}
        <div class="page">...</div>
    {/if}

    {#each Array.from({ length: totalPages }, (_, i) => i + 1) as page}
        {#if page === currentPage}
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <div class="page active-page" on:click={() => onChangePage(page)}>
                {page}
            </div>
        {:else if Math.abs(currentPage - page) < 3 || page === 1 || page === totalPages}
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <div class="page" on:click={() => onChangePage(page)}>{page}</div>
        {/if}
    {/each}

    {#if currentPage < totalPages - 2}
        <div class="page">...</div>
    {/if}

    {#if currentPage < totalPages}
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <div class="page" on:click={() => onChangePage(currentPage + 1)}>
            Next
        </div>
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <div class="page" on:click={() => onChangePage(totalPages)}>Last</div>
    {/if}
</div>

<style>
    .pagination {
        display: flex;
        justify-content: center;
        align-items: center;
        margin-top: 20px;
        margin-bottom: 15px;
    }

    .page {
        margin: 0 5px;
        padding: 5px 10px;
        cursor: pointer;
        background-color: #f0f0f0;
        border-radius: 5px;
        color: #333;
    }

    .active-page {
        font-weight: bold;
        background-color: #e9ab39;
        color: #fff;
    }

    .page:hover,
    .active-page:hover {
        scale: 1.1;
    }
</style>
