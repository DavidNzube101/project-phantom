<script>
	import { Bot, Eye, X, Sparkles } from 'lucide-svelte';
	import { Button } from '$lib/components/ui/button';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';

	let { searchQuery } = $props();
	let isRevealed = $state(false);
	let showModal = $state(false);
	let isGenerating = $state(false);

	const aiSummary = `Based on your search for "${searchQuery}", here's what I found:

**Key Insights:**
• Privacy tools have evolved significantly in 2024, with new focus on zero-knowledge architectures
• End-to-end encryption is becoming standard across messaging platforms
• Open-source alternatives are gaining mainstream adoption due to privacy concerns
• New regulations like GDPR 2.0 are pushing companies toward privacy-by-design

**Top Recommendations:**
1. **Signal** - Most secure messaging app with perfect forward secrecy
2. **Tor Browser** - Enhanced privacy browsing with improved speed
3. **ProtonMail** - Zero-access encryption for email communications
4. **Brave Browser** - Built-in ad blocking and privacy protection

**Recent Developments:**
The privacy landscape has shifted dramatically with major tech companies implementing new data collection policies. Users are increasingly seeking alternatives that prioritize privacy without sacrificing functionality.

**Privacy Score:** 9.2/10 - These tools offer excellent privacy protection with minimal data collection.`;

	function revealSummary() {
		isRevealed = true;
		isGenerating = true;
		
		setTimeout(() => {
			isGenerating = false;
			showModal = true;
		}, 1500);
	}

	function closeModal() {
		showModal = false;
	}
</script>

<Card class="border-purple-700/30 bg-gradient-to-r from-blue-900/20 to-purple-900/20 backdrop-blur-sm shadow-sm relative overflow-hidden">
	<CardContent class="p-6">
		<div class="flex items-start gap-4">
			<div class="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-blue-500 to-purple-500 flex-shrink-0">
				<Bot class="h-5 w-5 text-white" />
			</div>
			<div class="flex-1">
				<div class="flex items-center gap-2 mb-3">
					<h3 class="text-lg font-semibold text-white">AI Summary</h3>
					<Badge variant="secondary" class="bg-blue-500/20 text-blue-400 border-blue-500/30">
						<Sparkles class="h-3 w-3 mr-1" />
						Powered by AI
					</Badge>
				</div>
				
				{#if !isRevealed}
					<div class="relative">
						<div class="blur-sm select-none pointer-events-none text-purple-200 leading-relaxed mb-4">
							Based on your search for "{searchQuery}", here are the key insights about privacy tools and digital security. This comprehensive analysis covers the latest developments in privacy technology, recommended tools, and important considerations for protecting your digital footprint...
						</div>
						<div class="absolute inset-0 bg-gradient-to-r from-transparent via-purple-900/80 to-transparent"></div>
					</div>
					<Button
						onclick={revealSummary}
						class="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white"
					>
						<Eye class="h-4 w-4 mr-2" />
						Reveal AI Summary
					</Button>
				{:else if isGenerating}
					<div class="flex items-center gap-3 py-8">
						<div class="h-6 w-6 animate-spin rounded-full border-2 border-purple-500/30 border-t-purple-500"></div>
						<p class="text-purple-200">Generating AI summary...</p>
					</div>
				{:else}
					<div class="text-purple-200 leading-relaxed mb-4">
						<p class="line-clamp-3">
							Based on your search for "{searchQuery}", here are the key insights about privacy tools and digital security. This comprehensive analysis covers the latest developments...
						</p>
					</div>
					<Button
						onclick={() => showModal = true}
						variant="outline"
						class="border-purple-500 text-purple-300 hover:bg-purple-800/30 hover:text-white"
					>
						View Full Summary
					</Button>
				{/if}
			</div>
		</div>
	</CardContent>
</Card>

{#if showModal}
	<div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
		<div class="bg-purple-900/90 backdrop-blur-sm rounded-2xl shadow-2xl max-w-4xl w-full max-h-[80vh] overflow-hidden border border-purple-700/50">
			<div class="flex items-center justify-between p-6 border-b border-purple-700/50">
				<div class="flex items-center gap-3">
					<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-blue-500 to-purple-500">
						<Bot class="h-4 w-4 text-white" />
					</div>
					<h2 class="text-xl font-semibold text-white">AI Summary</h2>
					<Badge variant="secondary" class="bg-blue-500/20 text-blue-400 border-blue-500/30">
						<Sparkles class="h-3 w-3 mr-1" />
						Generated
					</Badge>
				</div>
				<Button
					variant="ghost"
					size="sm"
					onclick={closeModal}
					class="text-purple-400 hover:text-white hover:bg-purple-800/50"
				>
					<X class="h-5 w-5" />
				</Button>
			</div>
			
			<div class="p-6 overflow-y-auto max-h-[60vh]">
				<div class="prose prose-invert max-w-none text-purple-200">
					{@html aiSummary.replace(/\*\*(.*?)\*\*/g, '<strong class="text-white">$1</strong>').replace(/\n/g, '<br>').replace(/•/g, '&bull;')}
				</div>
			</div>
			
			<div class="flex items-center justify-between p-6 border-t border-purple-700/50 bg-purple-950/50">
				<p class="text-sm text-purple-400">
					Summary generated in 1.2 seconds • Privacy-focused analysis
				</p>
				<Button
					onclick={closeModal}
					class="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white"
				>
					Close
				</Button>
			</div>
		</div>
	</div>
{/if}

<style>
	.line-clamp-3 {
		display: -webkit-box;
		-webkit-line-clamp: 3;
		-webkit-box-orient: vertical;
		overflow: hidden;
	}
</style>
