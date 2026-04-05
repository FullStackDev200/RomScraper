<script lang="ts">
  import Spinner from "./Spinner.svelte";

  import { createEventDispatcher, onMount } from "svelte";

  import { ChooseDirectory, VimDownloadGame } from "../wailsjs/go/main/App.js";

  import { scraping } from "../wailsjs/go/models";

  export let selectedRoms: scraping.Rom[];
  let RomCount = 0;
  $: selectedRom = selectedRoms[RomCount];
  export let showDialog: Boolean;
  let dialog: HTMLDialogElement;
  let isLoading = false;

  onMount(async () => {
    window.addEventListener("keydown", (e) => {
      if (e.key == "Escape") {
        closeDialog();
      }
    });
  });

  const dispatch = createEventDispatcher();

  $: if (dialog) {
    if (showDialog && !dialog.open) dialog.showModal();
    else if (!showDialog && dialog.open) dialog.close();
  }

  function closeDialog() {
    dispatch("close");
  }

  async function pickDir() {
    const dir = await ChooseDirectory();
    if (dir) {
      VimDownloadGame(selectedRom, dir);
    } else {
    }
  }

  // TODO: Add RA rom validation
  let validatingPopup = false;

  async function validateRom() {
    validatingPopup = true;
    await new Promise((r) => setTimeout(r, 2000));
  }

  function goToDownlaodPage() {
    if (selectedRom.PageUrl) {
      // @ts-ignore
      window.runtime.BrowserOpenURL(selectedRom.PageUrl);
    }
  }

  function handleNext() {
    if (RomCount < selectedRoms.length - 1) {
      RomCount++;
    }
  }

  function handlePrev() {
    if (RomCount - 1 >= 0) {
      RomCount--;
    }
  }
</script>

<!-- TODO: make max sizes so that UI doesn't change that much -->
<dialog bind:this={dialog} class="noselect">
  <div class="game-info">
    <h2>{selectedRom.Title}</h2>
    <p><strong>Download Cover:</strong></p>
    <img src={selectedRom.CoverUrl} alt={selectedRom.Title} width="200" />
    <p><strong>Download Title:</strong> {selectedRom.Title}</p>
  </div>
  <div class="separator"></div>
  <div class="rom-info">
    <div class="rom-content">
      <p id="rom-info"><strong>Rom Info</strong></p>

      {#if !isLoading && selectedRoms}
        <p style="max-width: 10em">Name: {selectedRom.Title}</p>
        <p>Platform: {selectedRom.Platform}</p>
        <div id="choose-buttons">
          <button on:click={handlePrev}>&lt;&lt; Previous</button>
          <p>|</p>
          <button on:click={handleNext}>Next &gt;&gt;</button>
        </div>
        <div class="rom-select-info">
          <p style="margin-top: 1px;">{RomCount + 1}</p>
          <p>/</p>
          <p style="margin-bottom: -1px;">{selectedRoms.length}</p>
        </div>
        <div id="user-actions">
          <button on:click={goToDownlaodPage}>Go to Download Page</button>
          <button on:click={pickDir}>Download game</button>
          <button on:click={validateRom}>Validate Rom Hash</button>
        </div>
      {/if}
      {#if isLoading}
        <Spinner></Spinner>
      {/if}
      {#if !selectedRoms && !isLoading}
        <h1>Couldn't Load The Rom</h1>
      {/if}

      <button id="close-dialog" on:click={closeDialog}>Close</button>
    </div>
  </div>
</dialog>

<style>
  #close-dialog {
    margin-top: max(10px, 20%);
  }

  #user-actions {
    display: flex;
    align-items: center;
    flex-direction: column;
  }

  #user-actions button {
    border: 1px solid var(--accent-color);
    margin-bottom: 0.5em;
    width: 100%;
  }
  #choose-buttons {
    display: grid;
    grid-template-columns: auto auto auto;
    column-gap: 0.5em;
  }
</style>
