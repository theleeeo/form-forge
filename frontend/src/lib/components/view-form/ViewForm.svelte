<script lang="ts">
	import { Helper, Textarea, Button, Dropdown, DropdownItem } from 'flowbite-svelte';
	import { ChevronDownOutline } from 'flowbite-svelte-icons';
	import SortableList from './SortableList.svelte';
	import QuestionInput from './QuestionInput.svelte';
	import FloatingLableInputLocal from './FloatingLableInputLocal.svelte';
	import {
		CreateCheckboxQuestionParameters,
		CreateQuestionParameters,
		CreateRadioQuestionParameters,
		CreateRequest,
		CreateTextQuestionParameters,
		Form,
		Question,
		UpdateRequest
	} from '$lib/proto/form/v1/forms_pb.js';
	import { formClient } from '$lib/formClient.js';
	import { ConnectError } from '@connectrpc/connect';
	import { activeAlert } from '$lib/alert.ts';

	export let mode: 'view' | 'edit' | 'create' = 'view';
	export let form: Form | undefined = undefined;
	export let questions: Question[] = [];

	interface OrderableQuestion {
		order: number;
		type: 'text' | 'radio' | 'checkbox' | undefined;
		title: string;
		options: string[];
	}

	function convertFormQuestionsToOrderable(questions: Question[]): OrderableQuestion[] {
		let orderableQuestions: OrderableQuestion[] = [];

		questions.forEach((question, index) => {
			let orderableQuestion: OrderableQuestion = {
				order: index + 1,
				type: question.question.case,
				title: '',
				options: []
			};

			if (question.question.case === 'text') {
				orderableQuestion.title = question.question.value.title;
			} else if (question.question.case === 'radio') {
				orderableQuestion.title = question.question.value.title;
				orderableQuestion.options = question.question.value.options;
			} else if (question.question.case === 'checkbox') {
				orderableQuestion.title = question.question.value.title;
				orderableQuestion.options = question.question.value.options;
			} else {
				throw new Error('Unknown question type');
			}

			orderableQuestions.push(orderableQuestion);
		});

		return orderableQuestions;
	}

	function convertQuestionsToCreateParams(
		questions: OrderableQuestion[]
	): CreateQuestionParameters[] {
		let convertedQuestions: CreateQuestionParameters[] = [];

		questions.forEach((question) => {
			let questionParams = new CreateQuestionParameters();
			if (question.type === 'text') {
				questionParams.question = {
					case: 'text',
					value: new CreateTextQuestionParameters({
						title: question.title ?? ''
					})
				};
			} else if (question.type === 'radio') {
				questionParams.question = {
					case: 'radio',
					value: new CreateRadioQuestionParameters({
						title: question.title ?? '',
						options: question.options
					})
				};
			} else if (question.type === 'checkbox') {
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

			convertedQuestions.push(questionParams);
		});

		return convertedQuestions;
	}

	let questionDropdownOpen = false;
	const createQuestion = (questionType: 'text' | 'radio' | 'checkbox') => {
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

	const sortList = (ev: any) => {
		listQuestions = ev.detail;
	};

	const removeQuestion = (ev: any) => {
		listQuestions = listQuestions.filter((question) => question.order !== ev.detail.order);
	};

	function submitForm() {
		let questions = convertQuestionsToCreateParams(listQuestions);

		if (mode === 'edit') {
			formClient
				.update(
					new UpdateRequest({
						baseId: form?.baseId,
						newForm: new CreateRequest({
							title: title,
							description: description,
							questions: questions
						})
					})
				)
				.then(() => {
					activeAlert.set({ type: 'success', message: 'Form updated' });
				})
				.catch((error) => {
					activeAlert.set({ type: 'error', message: ConnectError.from(error).message });
				});
			return;
		} else {
			formClient
				.create(
					new CreateRequest({
						title: title,
						description: description,
						questions: questions
					})
				)
				.then(() => {
					activeAlert.set({ type: 'success', message: 'Form created' });
				})
				.catch((error) => {
					activeAlert.set({ type: 'error', message: ConnectError.from(error).message });
				});
		}
	}

	let title: string = '';
	let description: string = '';
	let listQuestions: OrderableQuestion[] = [];

	let disabled: boolean = false;

	if (mode === 'view') {
		if (form === undefined) {
			throw new Error('Form is required in view mode');
		}

		title = form.title;
		description = form.description;
		listQuestions = convertFormQuestionsToOrderable(questions);

		disabled = true;
	}

	if (mode === 'edit') {
		if (form === undefined) {
			throw new Error('Form is required in edit mode');
		}

		title = form.title;
		description = form.description;
		listQuestions = convertFormQuestionsToOrderable(questions);
	}
</script>

<div class="input-field">
	<FloatingLableInputLocal
		id="title-input"
		type="text"
		classDiv="mb-4"
		classInput={disabled ? 'cursor-not-allowed' : ''}
		bind:value={title}
		{disabled}
	>
		Title
	</FloatingLableInputLocal>

	<Textarea
		id="descr-input"
		rows={4}
		placeholder="Add a description..."
		bind:value={description}
		{disabled}
	/>
	<Helper class="text-sm text-gray-500"
		>The description will only be shown on this admin page</Helper
	>
</div>

<div class="input-field">
	<h3 class="mb-4 text-lg font-bold text-gray-800 dark:text-gray-200">Questions:</h3>

	<SortableList class="mb-4" list={listQuestions} key="order" on:sort={sortList} let:item>
		<QuestionInput question={item} on:remove={removeQuestion} {disabled} />
	</SortableList>

	{#if disabled}
		<Button
			class="mt-4"
			on:click={() => {
				disabled = false;
				mode = 'edit';
			}}>Edit</Button
		>
	{:else}
		<Button class="ml-auto"
			>Add Question<ChevronDownOutline class="ms-2 h-6 w-6 text-white" /></Button
		>
		<Dropdown bind:open={questionDropdownOpen}>
			<DropdownItem on:click={() => createQuestion('text')}>text</DropdownItem>
			<DropdownItem on:click={() => createQuestion('radio')}>radio</DropdownItem>
			<DropdownItem on:click={() => createQuestion('checkbox')}>checkbox</DropdownItem>
		</Dropdown>

		<Button class="mt-4" on:click={submitForm}>{mode === 'create' ? 'Create' : 'Update'}</Button>
	{/if}
</div>

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
