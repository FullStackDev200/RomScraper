<script lang="ts">
  import { groupType } from "./types";
  export let selected: groupType[] = [];

  function toggle(option: groupType) {
    if (selected.includes(option)) {
      selected = selected.filter((o) => o !== option);
    } else {
      selected = [...selected.filter((e) => e !== groupType.ByNone), option];
    }
    if (selected.includes(groupType.ByNone) || selected.length === 0) {
      selected = [groupType.ByNone];
    }
  }

  // TODO: make None be automatically unselected when another thing is selected
</script>

<div class="group-selector">
  <span class="label">Group by</span>
  <div class="chips">
    {#each Object.values(groupType) as option}
      <button
        class="chip"
        class:active={selected.includes(option)}
        on:click={() => toggle(option)}
      >
        {option}
      </button>
    {/each}
  </div>
</div>

<style>
  .group-selector {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .label {
    font-size: 11px;
    font-weight: 600;
    letter-spacing: 0.1em;
    text-transform: uppercase;
    white-space: nowrap;
  }

  .chips {
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
  }

  .chip {
    padding: 5px 14px;
    border-radius: 999px;
    border: 1.5px solid #333;
    background: transparent;
    font-size: 12px;
    font-weight: 500;
    letter-spacing: 0.04em;
    cursor: pointer;
    transition:
      background 0.15s,
      color 0.15s,
      border-color 0.15s,
      box-shadow 0.15s;
  }

  .chip:hover {
    border-color: #888;
    color: #fff;
  }

  .chip.active {
    background: #fff;
    border-color: #fff;
    color: #111;
    box-shadow: 0 0 10px rgba(255, 255, 255, 0.15);
  }
</style>
