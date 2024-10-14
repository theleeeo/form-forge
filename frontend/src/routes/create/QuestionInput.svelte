<script lang="ts">
	// import { formatDateDifference } from '@grpc/grpc-js/build/src/deadline';
	import {
		Label,
		Input,
		Helper,
		Textarea,
		Button,
		Dropdown,
		DropdownItem,
		FloatingLabelInput,
		Listgroup,
		InputAddon
	} from 'flowbite-svelte';
	import { TrashBinOutline } from 'flowbite-svelte-icons';
	import { createEventDispatcher } from 'svelte';

	// TODO: Remove duplicate from create/+page.svelte
	interface Question {
		order: number;
		type: 'Text' | 'Radio' | 'Checkbox';
		title: string | undefined;
		options: string[];
	}

	export let question: Question;

	// TODO: Required?
	let newOption = '';

	const dispatch = createEventDispatcher();

	function removeQuestion() {
		console.log('Removing question', question.order);
		dispatch('remove', { order: question.order });
	}
</script>

{#if question.type === 'Text'}
	<fieldset class="relative mb-2 rounded-md border border-gray-300 pt-1 dark:border-gray-500">
		<legend class="mx-2 px-2 dark:text-gray-100">{question.type}</legend>
		<button
			class="absolute right-0 top-0 mr-1 mt-1 cursor-pointer text-red-500 dark:text-red-400"
			on:click={removeQuestion}
		>
			<TrashBinOutline />
		</button>
		<div class="mb-2 flex">
			<Label class="mx-2" for="{question.order}-title">Title:</Label>
			<input
				id="{question.order}-title"
				class="rounded-t-md border-0 border-b border-b-gray-500 pl-1 dark:bg-gray-600 dark:text-gray-200"
				type="text"
				placeholder="..."
				bind:value={question.title}
			/>
		</div>
	</fieldset>
{:else if question.type === 'Radio'}
	<fieldset class="relative mb-2 rounded-md border border-gray-300 pt-1 dark:border-gray-500">
		<legend class="mx-2 px-2 dark:text-gray-100">{question.type}</legend>
		<button
			class="absolute right-0 top-0 mr-1 mt-1 cursor-pointer text-red-500 dark:text-red-400"
			on:click={removeQuestion}
		>
			<TrashBinOutline />
		</button>
		<div class="flex pb-4">
			<Label class="mx-2" for="{question.order}-title">Title:</Label>
			<input
				id="{question.order}-title"
				class="rounded-t-md border-b border-b-gray-500 pl-1 dark:bg-gray-600 dark:text-gray-200"
				type="text"
				placeholder="..."
				bind:value={question.title}
			/>
		</div>

		{#if question.options}
			{#each question.options as _, i}
				<div class="mb-4 ml-2 flex items-center">
					<Label class="mx-2" for="{question.order}-option-next">Option:</Label>
					<input
						class="rounded-t-md border-b border-b-gray-500 pl-1 dark:bg-gray-600 dark:text-gray-200"
						type="text"
						id="{question.order}-option-{i}"
						name="quesion-{question.order}-option-{i}"
						placeholder="..."
						bind:value={question.options[i]}
						on:change={() => {
							if (question.options[i] === '') {
								question.options = question.options.filter((_, index) => index !== i);
							}
						}}
					/>
				</div>
			{/each}
		{/if}

		<Label class="ml-4" for="{question.order}-option-next">New Option:</Label>
		<input
			class="mb-4 ml-4 rounded-t-md border-b border-b-gray-500 pl-1 dark:bg-gray-600 dark:text-gray-300"
			type="text"
			id="{question.order}-option-next"
			name="quesion-{question.order}-option-next"
			placeholder="..."
			bind:value={newOption}
			on:change={() => {
				if (question.options === null) {
					question.options = [];
				}

				question.options = [...question.options, newOption];
				newOption = '';
			}}
		/>
	</fieldset>
{:else if question.type === 'Checkbox'}
	<fieldset class="relative mb-2 rounded-md border border-gray-300 pt-1 dark:border-gray-500">
		<legend class="mx-2 px-2 dark:text-gray-100">{question.type}</legend>
		<button
			class="absolute right-0 top-0 mr-1 mt-1 cursor-pointer text-red-500 dark:text-red-400"
			on:click={removeQuestion}
		>
			<TrashBinOutline />
		</button>
		<div class="flex pb-4">
			<Label class="mx-2" for="{question.order}-title">Title:</Label>
			<input
				id="{question.order}-title"
				class="rounded-t-md border-b border-b-gray-500 pl-1 dark:bg-gray-600 dark:text-gray-200"
				type="text"
				placeholder="..."
				bind:value={question.title}
			/>
		</div>

		{#if question.options}
			{#each question.options as _, i}
				<div class="mb-4 ml-2 flex items-center">
					<Label class="mx-2" for="{question.order}-option-{i}">Option:</Label>
					<input
						class="rounded-t-md border-b border-b-gray-500 pl-1 dark:bg-gray-600 dark:text-gray-200"
						type="text"
						id="{question.order}-option-{i}"
						name="quesion-{question.order}-option-{i}"
						placeholder="..."
						bind:value={question.options[i]}
						on:change={() => {
							if (question.options[i] === '') {
								question.options = question.options.filter((_, index) => index !== i);
							}
						}}
					/>
				</div>
			{/each}
		{/if}

		<Label class="ml-4" for="{question.order}-option-next">New Option:</Label>
		<input
			class="mb-4 ml-4 rounded-t-md border-b border-b-gray-500 pl-1 dark:bg-gray-600 dark:text-gray-300"
			type="text"
			id="{question.order}-option-next"
			name="quesion-{question.order}-option-next"
			placeholder="..."
			bind:value={newOption}
			on:change={() => {
				question.options = [...question.options, newOption];
				newOption = '';
			}}
		/>
	</fieldset>
{/if}
