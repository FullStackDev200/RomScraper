<script lang="ts">
  import search from "./assets/icons/search-icon.svg";
  import settings from "./assets/icons/settings-svgrepo-com.svg";

  import Modal from "./Modal.svelte";
  import Settings from "./Settings.svelte";

  import { GetGamesByName, GetGameCoverUrl } from "../wailsjs/go/main/App.js";

  let selectedGame = null;
  let showDialog = false;
  let showSettings = false;

  /**
   * Handle a game object from TheGamesDB scraping
   * @param {Object} game - A game object containing data from scraping\thegamesdb.go
   */
  function openDialog(game) {
    selectedGame = game;
    showDialog = true;
  }

  let games = [];
  let text = "";
  let loading = false;
  let error = "";

  async function handleSubmit() {
    const query = text.trim();
    if (!query) return;

    text = "";
    games = [];
    error = "";
    loading = true;

    try {
      const fetchedGames = await GetGamesByName(query);

      if (fetchedGames?.length) {
        const covers = await Promise.all(
          fetchedGames.map((g) => GetGameCoverUrl(g.Id)),
        );
        games = fetchedGames.map((g, i) => ({ ...g, CoverUrl: covers[i] }));
      } else {
        error = "No games found.";
      }
    } catch (err) {
      console.error("Error fetching games:", err);
      error = "Failed to load games.";
    } finally {
      loading = false;
    }
  }
</script>

<link href="./style.css" />

<main>
  <div class="topnav">
    <form on:submit|preventDefault={handleSubmit} class="searchbar">
      <input type="text" bind:value={text} placeholder="Type something..." />
      <button type="submit">
        <img src={search} width="20" alt="search-icon" class="search icon" />
      </button>
    </form>

    <button
      class="settings-button"
      on:click={() => {
        showSettings = true;
      }}
      ><img src={settings} alt="settings icons" class="settings icon" /></button
    >
  </div>

  <div class="roms">
    {#each games as game}
      <button class="rom" on:click={() => openDialog(game)}>
        {#if game.CoverUrl}
          <img src={game.CoverUrl} alt={game.Title} />
        {:else}
          <p>Loading cover...</p>
        {/if}
        <p>{game.Title}</p>
      </button>
    {/each}
    {#if showDialog}
      <Modal
        bind:selectedGame
        bind:showDialog
        on:close={() => (showDialog = false)}
      ></Modal>
    {/if}
    {#if showSettings}
      <Settings on:close={() => (showSettings = false)}></Settings>
    {/if}
  </div>
</main>
