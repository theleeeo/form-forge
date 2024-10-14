<script lang="ts">
	import { formClient } from '$lib/formClient.js';
	import { Form, Question, TextQuestion } from '$lib/proto/form/v1/forms_pb.js';
	import { AccordionItem, Accordion, Card } from 'flowbite-svelte';

	export let form: Form;
	export let questions: Array<Question>;

	console.log(questions);
</script>

<h2 class="text-xl font-bold text-gray-900 dark:text-gray-100">
	{form.title}
</h2>
<p class="mb-4 text-sm">
	{form.description}
</p>
<h2 class="text-xl font-bold text-gray-900 dark:text-gray-100">Questions:</h2>
<ul>
	{#each questions as question}
		{#if question.question.case === 'text'}
			<li>
				{question.question.value.title}
			</li>
		{:else if question.question.case === 'radio'}
			<li>
				<ul>
					{question.question.value.title}
					{#each question.question.value.options as option}
						<li>
							{option}
						</li>
					{/each}
				</ul>
			</li>
		{:else if question.question.case === 'checkbox'}
			<li>
				<ul>
					{question.question.value.title}
					{#each question.question.value.options as option}
						<li>
							{option}
						</li>
					{/each}
				</ul>
			</li>
		{/if}
	{/each}
</ul>

<style>
	ul {
		list-style-type: none;
		padding: 0;
	}

	li {
		margin: 0.5rem 0;
		padding: 0.5rem;
		background-color: #f9f9f9;
		border: 1px solid #ddd;
		border-radius: 4px;
	}

	li ul {
		margin-top: 0.5rem;
	}

	li ul li {
		background-color: #fff;
		border: none;
		padding-left: 1rem;
	}
</style>
