<script>
	import { Brain, CheckCircle, Clock, MapPin, User } from 'lucide-svelte';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';

	let { data } = $props();
	
	const answers = {
		'bitcoin': {
			title: 'What is Bitcoin?',
			answer: 'Bitcoin is a decentralized digital currency that operates without a central bank or single administrator. It uses blockchain technology to enable peer-to-peer transactions and is secured by cryptographic proof.',
			source: 'Cryptocurrency Encyclopedia',
			icon: CheckCircle
		},
		'encrypt files': {
			title: 'How to encrypt files',
			answer: 'File encryption converts your data into unreadable code using algorithms. You can use built-in tools like BitLocker (Windows) or FileVault (Mac), or third-party software like VeraCrypt for cross-platform encryption.',
			source: 'Security Best Practices',
			icon: CheckCircle
		}
	};

	function getAnswer(query) {
		const lowerQuery = query.toLowerCase();
		for (const [key, value] of Object.entries(answers)) {
			if (lowerQuery.includes(key)) {
				return value;
			}
		}
		return {
			title: `About "${data.query}"`,
			answer: `Based on your search for "${data.query}", here's what we found: This appears to be a ${data.type} query. Our privacy-focused search provides comprehensive results while protecting your data.`,
			source: 'Breeze Knowledge Base',
			icon: Brain
		};
	}

	const answer = getAnswer(data.fullQuery);
	const IconComponent = answer.icon;
</script>

<Card class="border-purple-700/30 bg-gradient-to-r from-purple-900/30 to-blue-900/20 backdrop-blur-sm shadow-sm">
	<CardContent class="p-6">
		<div class="flex items-start gap-4">
			<div class="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-purple-500 to-blue-500 flex-shrink-0">
				<IconComponent class="h-5 w-5 text-white" />
			</div>
			<div class="flex-1">
				<div class="flex items-center gap-2 mb-3">
					<h3 class="text-lg font-semibold text-white">{answer.title}</h3>
					<Badge variant="secondary" class="bg-purple-500/20 text-purple-400 border-purple-500/30">
						<Brain class="h-3 w-3 mr-1" />
						Direct Answer
					</Badge>
				</div>
				
				<p class="text-purple-200 leading-relaxed mb-4">
					{answer.answer}
				</p>
				
				<div class="flex items-center gap-2 text-sm text-purple-400">
					<CheckCircle class="h-4 w-4" />
					<span>Source: {answer.source}</span>
				</div>
			</div>
		</div>
	</CardContent>
</Card>
