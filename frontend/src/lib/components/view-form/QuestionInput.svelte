<script lang="ts">
	import { Value } from '@bufbuild/protobuf';
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
		InputAddon,
		Card
	} from 'flowbite-svelte';
	import { TrashBinOutline } from 'flowbite-svelte-icons';
	import { createEventDispatcher } from 'svelte';
	import FloatingLableInputLocal from './FloatingLableInputLocal.svelte';

	// TODO: Remove duplicate from create/+page.svelte
	interface Question {
		order: number;
		type: 'text' | 'radio' | 'checkbox';
		title: string | undefined;
		options: string[];
	}

	export let question: Question;
	export let disabled = false;

	// TODO: Required?
	let newOption = '';

	const dispatch = createEventDispatcher();

	function removeQuestion() {
		dispatch('remove', { order: question.order });
	}
</script>

<fieldset class="relative mb-2 rounded-md border border-gray-300 p-4 pt-2 dark:border-gray-500">
	<legend class="mx-2 px-2 dark:text-gray-100"
		>{question.type.charAt(0).toUpperCase() + question.type.slice(1)}</legend
	>
	{#if !disabled}
		<button
			class="absolute right-0 top-0 mr-1 mt-1 cursor-pointer text-red-500 dark:text-red-400"
			on:click={removeQuestion}
		>
			<TrashBinOutline />
		</button>
	{/if}

	<div class="flex items-center">
		<Label class="mr-2" for="{question.order}-title">Title:</Label>
		<FloatingLableInputLocal
			id="{question.order}-title"
			placeholder="..."
			bind:value={question.title}
			type="text"
			classDiv="w-full"
			{disabled}
		/>
	</div>
	{#if question.type === 'radio' || question.type === 'checkbox'}
		{#each question.options as _, i}
			<div class="m-2 ml-2 flex items-center">
				<Label class="mr-2" for="{question.order}-option-{i}">Option:</Label>
				<FloatingLableInputLocal
					id="{question.order}-option-{i}"
					name="quesion-{question.order}-option-{i}"
					placeholder="..."
					type="text"
					classInput="pb-2"
					bind:value={question.options[i]}
					{disabled}
					on:change={() => {
						if (question.options[i] === '') {
							question.options = question.options.filter((_, index) => index !== i);
						}
					}}
				/>
			</div>
		{/each}

		{#if !disabled}
			<div class="m-2 ml-2 flex items-center">
				<Label class="mr-2" for="{question.order}-option-next">New Option:</Label>
				<FloatingLableInputLocal
					id="{question.order}-option-next"
					name="quesion-{question.order}-option-next"
					placeholder="..."
					type="text"
					classInput="pb-2"
					bind:value={newOption}
					{disabled}
					on:change={() => {
						question.options = [...question.options, newOption];
						newOption = '';
					}}
				/>
			</div>
		{/if}
	{/if}
</fieldset>
