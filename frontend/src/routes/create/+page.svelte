<script lang="ts">
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
		Alert
	} from 'flowbite-svelte';
	import { ChevronDownOutline } from 'flowbite-svelte-icons';
	import SortableList from '$lib/SortableList.svelte';
	import QuestionInput from './QuestionInput.svelte';
	import FloatingLableInputLocal from './FloatingLableInputLocal.svelte';
	import {
		CreateCheckboxQuestionParameters,
		CreateQuestionParameters,
		CreateRadioQuestionParameters,
		CreateRequest,
		CreateTextQuestionParameters
	} from '$lib/proto/form/v1/forms_pb.js';
	import { formClient } from '$lib/formClient.js';
	import { ConnectError } from '@connectrpc/connect';

	interface Question {
		order: number;
		type: 'Text' | 'Radio' | 'Checkbox';
		title: string;
		options: string[];
	}

	let questionDropdownOpen = false;
	const createQuestion = (questionType: 'Text' | 'Radio' | 'Checkbox') => {
		questionDropdownOpen = false;

		listQuestions = [
			...listQuestions,
			{
				order: listQuestions.length + 1,
				type: questionType,
				title: '',
				options: []
			}
		];
	};

	let listQuestions: Question[] = [
		{
			order: 1,
			type: 'Radio',
			title: 'What is your favorite color?',
			options: ['Red', 'Green', 'Blue']
		},
		{
			order: 2,
			type: 'Text',
			title: 'What is your favorite color?',
			options: []
		},
		{
			order: 3,
			type: 'Checkbox',
			title: 'What is your favorite color?',
			options: ['Red', 'Green', 'Blue']
		}
	];

	const sortList = (ev: any) => {
		listQuestions = ev.detail;
	};

	const removeQuestion = (ev: any) => {
		listQuestions = listQuestions.filter((question) => question.order !== ev.detail.order);
	};

	function submitForm() {
		let questions: CreateQuestionParameters[] = [];

		listQuestions.forEach((question) => {
			let questionParams = new CreateQuestionParameters();
			if (question.type === 'Text') {
				questionParams.question = {
					case: 'text',
					value: new CreateTextQuestionParameters({
						title: question.title ?? ''
					})
				};
			} else if (question.type === 'Radio') {
				questionParams.question = {
					case: 'radio',
					value: new CreateRadioQuestionParameters({
						title: question.title ?? '',
						options: question.options
					})
				};
			} else if (question.type === 'Checkbox') {
				questionParams.question = {
					case: 'checkbox',
					value: new CreateCheckboxQuestionParameters({
						title: question.title ?? '',
						options: question.options
					})
				};
			} else {
				throw new Error('Unknown question type');
			}

			questions.push(questionParams);
		});

		console.log('Submitting form', title, description, questions);

		formClient
			.create(
				new CreateRequest({
					title: title,
					description: description,
					questions: questions
				})
			)
			.then((response) => {
				console.log('Form created', response);
			})
			.catch((error) => {
				showAlert(ConnectError.from(error).message);
				console.error('Error creating form', error);
			});
	}

	let title: string = '';
	let description: string = '';

	let alert: string = '';
	const showAlert = (message: string) => {
		alert = message;
		setTimeout(() => {
			alert = '';
		}, 5000);
	};
</script>

<div class="input-field">
	<!-- <Label class="m-2 text-md w-fit" for="title-input">Title</Label>
	<Input class="bg-gray-100 mb-4" id="title-input" type="text" /> -->
	<FloatingLableInputLocal id="title-input" type="text" classDiv="mb-4" bind:value={title}>
		Title
	</FloatingLableInputLocal>

	<!-- <Label class="m-2 text-md w-fit" for="descr-input">Description</Label> -->
	<Textarea id="descr-input" rows={4} placeholder="Add a description..." bind:value={description} />
	<Helper class="text-sm text-gray-500"
		>The description will only be shown on this admin page</Helper
	>
</div>

<div class="input-field">
	<h3 class="mb-4 text-lg font-bold text-gray-800 dark:text-gray-200">Questions:</h3>

	<SortableList class="mb-4" list={listQuestions} key="order" on:sort={sortList} let:item>
		<QuestionInput question={item} on:remove={removeQuestion} />
	</SortableList>

	<Button class="ml-auto">Add Question<ChevronDownOutline class="ms-2 h-6 w-6 text-white" /></Button
	>
	<Dropdown bind:open={questionDropdownOpen}>
		<DropdownItem on:click={() => createQuestion('Text')}>Text</DropdownItem>
		<DropdownItem on:click={() => createQuestion('Radio')}>Radio</DropdownItem>
		<DropdownItem on:click={() => createQuestion('Checkbox')}>Checkbox</DropdownItem>
	</Dropdown>

	<Button class="mt-4" on:click={submitForm}>Submit</Button>
</div>

{#if alert}
	<Alert color="red" class="fixed top-0 m-4">
		{alert}
	</Alert>
{/if}

<style>
	.input-field {
		margin: 1em;
		/* border: 1px solid black; */

		min-width: 50%;

		padding: 1em;
		border-radius: 0.5em;

		display: flex;
		flex-direction: column;
	}
</style>
