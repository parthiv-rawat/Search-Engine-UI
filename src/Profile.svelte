<!-- Profile.svelte -->
<script>
    let name = "Parthiv";
    let age = 30;
    let email = "parthiv@example.com";
    let bio = "Hello, I'm Parthiv Rawat. I love programming and Svelte!";

    let isEditing = false;

    function saveChanges() {
        isEditing = false;
    }

    function cancelEditing() {
        isEditing = false;
    }
</script>

<header>
    <a href="/">
        <!-- {#if isVisible} -->
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <img src="images/image.png" alt="Logo" class="logo" />
        <!-- {/if} -->
        <h1>FlyHigh Edutech Search Engine</h1>
    </a>
    <button on:click>HomePage</button>
</header>

<main>
    <h2>Welcome to My Profile</h2>
    <div class="profile-image" />
    <section>
        <h3>Personal Information</h3>
        {#if isEditing}
            <div>
                <label>
                    Name:
                    <input bind:value={name} />
                </label>
            </div>
            <div>
                <label>
                    Age:
                    <input type="number" bind:value={age} />
                </label>
            </div>
            <div>
                <label>
                    Email:
                    <input type="email" bind:value={email} />
                </label>
            </div>
            <div>
                <button on:click={saveChanges}>Save</button>
                <button on:click={cancelEditing}>Cancel</button>
            </div>
        {:else}
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <div class="editable" on:click={() => (isEditing = true)}>
                <p>
                    <strong>Name:</strong>
                    {name}
                </p>
                <p>
                    <strong>Age:</strong>
                    {age}
                </p>
                <p>
                    <strong>Email:</strong>
                    {email}
                </p>
            </div>
        {/if}
    </section>
    <section>
        <h3>About Me</h3>
        <p>{bio}</p>
    </section>
</main>

<style>
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
        /* background-color: #fff; */
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
        /* align-items: center; */
        justify-content: center;
        background-color: #bbbbbe;
        min-width: 1310px;
        margin: 0 auto;
        padding: 20px;
        font-family: Arial, sans-serif;
        color: #333;
        border-radius: 10px;
        /* background-image: url("bg.avif");
        background-size: cover;
        background-position: center; */
    }

    h2 {
        text-align: center;
        font-size: 2rem;
        margin-bottom: 20px;
        color: #007bff; /* You can choose a color of your preference */
    }

    .editable {
        cursor: pointer;
    }

    section {
        margin-bottom: 30px;
    }

    h3 {
        font-size: 1.5rem;
        margin-bottom: 10px;
    }

    p {
        margin: 5px 0;
        line-height: 1.6;
    }

    strong {
        font-weight: bold;
    }

    input {
        width: 100%;
        padding: 5px;
        border: 1px solid #ccc;
        border-radius: 5px;
        margin-bottom: 10px;
    }

    button {
        background-color: #007bff;
        color: #fff;
        border: none;
        padding: 8px 16px;
        border-radius: 5px;
        cursor: pointer;
        margin-right: 10px;
    }

    .profile-image {
        width: 150px;
        height: 150px;
        border-radius: 50%;
        margin: 0 auto 20px;
        display: block;
        display: flex;
        justify-content: center;
        align-items: center;
        /* background-color: #ccc; You can add a background image or use your own image */
        background-image: url("avatar.png");
        background-position: center;
        background-size: auto;
    }
</style>
