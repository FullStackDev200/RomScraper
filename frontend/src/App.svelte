<script lang="ts">
  import search from "./assets/icons/search-icon.svg";
  import settings from "./assets/icons/settings-svgrepo-com.svg";

  import { groupType } from "./types";

  import Modal from "./Modal.svelte";
  import Settings from "./Settings.svelte";
  import Dropdown from "./Dropdown.svelte";

  import {
    GetGamesByName,
    GetGameCoverUrl,
    PlatformToString,
  } from "../wailsjs/go/main/App.js";

  import { scraping } from "../wailsjs/go/models";

  let selectedRoms: scraping.Rom[];
  let showDialog = false;
  let showSettings = false;

  function openDialog(roms: scraping.Rom[]) {
    selectedRoms = roms;
    showDialog = true;
  }

  let roms: scraping.Rom[] = [];
  let romgroups: Record<string, scraping.Rom[]>;
  let text = "";

  let selectedGroupTypes: groupType[] = [groupType.ByNone];

  // TODO:  Add caching
  async function handleSubmit() {
    const query = text.trim();
    if (!query) return;

    text = "";
    roms = [];

    const fetchedGames = await GetGamesByName(query);

    if (fetchedGames?.length) {
      const covers = await Promise.all(
        fetchedGames.map((g) => GetGameCoverUrl(g)),
      );
      roms = await Promise.all(
        fetchedGames.map(async (g, i) => {
          return new scraping.Rom({
            ...g,
            CoverUrl: covers[i],
            Platform: await PlatformToString(g.Platform),
          });
        }),
      );
    }
  }

  // TODO: Make platform display actual platform
  function groupBy(
    roms: scraping.Rom[],
    groupTypes: groupType[],
  ): Record<string, scraping.Rom[]> {
    const result: Record<string, scraping.Rom[]> = {};

    if (groupTypes.includes(groupType.ByNone)) {
      for (const rom of roms) {
        result[rom.Title] = [rom];
      }
      return result;
    }

    for (const rom of roms) {
      const key = groupTypes
        .map((grouptype) => {
          const value = rom[grouptype];
          return value;
        })
        .join("|");

      if (!result[key]) result[key] = [];
      result[key].push(rom);
    }

    return result;
  }

  $: romgroups = groupBy(roms, selectedGroupTypes);
  $: console.log(selectedGroupTypes);

  // TODO: Add a store to download multiple files at the same time
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

  <Dropdown bind:selected={selectedGroupTypes} />

  <!-- TODO: Make animation that will cycle between several covers-->
  <!-- remove unneccesary endings in rom file names  -->
  <div class="roms">
    {#each Object.entries(romgroups ?? {}) as [groupName, roms]}
      <button class="rom" on:click={() => openDialog(roms)}>
        {#if roms[0].CoverUrl}
          <img src={roms[0].CoverUrl} alt={roms[0].Title} />
        {:else}
          <p>Loading cover...</p>
        {/if}
        <p>{groupName}</p>
      </button>
    {/each}
    {#if showDialog}
      <Modal
        bind:selectedRoms
        bind:showDialog
        on:close={() => (showDialog = false)}
      ></Modal>
    {/if}
  </div>

  {#if showSettings}
    <Settings on:close={() => (showSettings = false)}></Settings>
  {/if}
</main>
