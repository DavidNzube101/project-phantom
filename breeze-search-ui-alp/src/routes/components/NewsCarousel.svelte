<script>
	import { ChevronLeft, ChevronRight, ExternalLink, Clock, TrendingUp } from 'lucide-svelte';
	import { Button } from '$lib/components/ui/button';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';

	let scrollContainer = $state();
	
	const newsItems = [
		{
			title: 'New Privacy Regulations Take Effect in EU',
			source: 'Privacy News',
			time: '2h ago',
			category: 'Policy',
			image: '/placeholder.svg?height=120&width=200',
			url: 'https://privacy-news.com/eu-regulations'
		},
		{
			title: 'Major Data Breach Affects 50M Users',
			source: 'Security Today',
			time: '4h ago',
			category: 'Security',
			image: '/placeholder.svg?height=120&width=200',
			url: 'https://security-today.com/data-breach'
		},
		{
			title: 'Signal Introduces New Privacy Features',
			source: 'Tech Privacy',
			time: '6h ago',
			category: 'Apps',
			image: '/placeholder.svg?height=120&width=200',
			url: 'https://tech-privacy.com/signal-features'
		},
		{
			title: 'Apple Enhances Safari Privacy Controls',
			source: 'Apple Insider',
			time: '8h ago',
			category: 'Browsers',
			image: '/placeholder.svg?height=120&width=200',
			url: 'https://apple-insider.com/safari-privacy'
		},
		{
			title: 'VPN Usage Surges 300% This Year',
			source: 'Digital Trends',
			time: '12h ago',
			category: 'VPN',
			image: '/placeholder.svg?height=120&width=200',
			url: 'https://digital-trends.com/vpn-usage'
		},
		{
			title: 'Google Faces New Privacy Lawsuit',
			source: 'Legal Tech',
			time: '1d ago',
			category: 'Legal',
			image: '/placeholder.svg?height=120&width=200',
			url: 'https://legal-tech.com/google-lawsuit'
		}
	];

	function scrollLeft() {
		scrollContainer.scrollBy({ left: -300, behavior: 'smooth' });
	}

	function scrollRight() {
		scrollContainer.scrollBy({ left: 300, behavior: 'smooth' });
	}

	function getCategoryColor(category) {
		const colors = {
			'Policy': 'bg-blue-500/20 text-blue-400 border-blue-500/30',
			'Security': 'bg-red-500/20 text-red-400 border-red-500/30',
			'Apps': 'bg-green-500/20 text-green-400 border-green-500/30',
			'Browsers': 'bg-purple-500/20 text-purple-400 border-purple-500/30',
			'VPN': 'bg-orange-500/20 text-orange-400 border-orange-500/30',
			'Legal': 'bg-gray-500/20 text-gray-400 border-gray-500/30'
		};
		return colors[category] || colors['Policy'];
	}
</script>

<Card class="border-purple-700/30 bg-purple-900/20 backdrop-blur-sm shadow-sm">
	<CardContent class="p-6">
		<div class="flex items-center justify-between mb-6">
			<div class="flex items-center gap-3">
				<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-orange-500 to-red-500">
					<TrendingUp class="h-4 w-4 text-white" />
				</div>
				<h3 class="text-lg font-semibold text-white">Latest Privacy News</h3>
			</div>
			<div class="flex gap-2">
				<Button
					variant="outline"
					size="sm"
					onclick={scrollLeft}
					class="h-8 w-8 p-0 border-purple-600 text-purple-300 hover:bg-purple-800/30 hover:text-white"
				>
					<ChevronLeft class="h-4 w-4" />
				</Button>
				<Button
					variant="outline"
					size="sm"
					onclick={scrollRight}
					class="h-8 w-8 p-0 border-purple-600 text-purple-300 hover:bg-purple-800/30 hover:text-white"
				>
					<ChevronRight class="h-4 w-4" />
				</Button>
			</div>
		</div>
		
		<div 
			bind:this={scrollContainer}
			class="flex gap-4 overflow-x-auto scrollbar-hide pb-2"
			style="scroll-snap-type: x mandatory;"
		>
			{#each newsItems as item}
				<div class="flex-shrink-0 w-72" style="scroll-snap-align: start;">
					<Card class="border-purple-700/30 bg-purple-900/30 hover:shadow-md transition-all duration-200 cursor-pointer group">
						<CardContent class="p-0">
							<div class="relative">
								<img
									src={item.image || "/placeholder.svg"}
									alt={item.title}
									class="w-full h-32 object-cover rounded-t-lg"
								/>
								<div class="absolute top-3 left-3">
									<Badge variant="secondary" class={getCategoryColor(item.category)}>
										{item.category}
									</Badge>
								</div>
							</div>
							<div class="p-4">
								<h4 class="font-semibold text-white mb-2 line-clamp-2 group-hover:text-purple-400 transition-colors">
									{item.title}
								</h4>
								<div class="flex items-center justify-between text-sm text-purple-400">
									<span>{item.source}</span>
									<div class="flex items-center gap-1">
										<Clock class="h-3 w-3" />
										<span>{item.time}</span>
									</div>
								</div>
								<a
									href={item.url}
									class="inline-flex items-center gap-1 text-sm text-purple-400 hover:text-purple-300 mt-2 transition-colors"
								>
									Read more
									<ExternalLink class="h-3 w-3" />
								</a>
							</div>
						</CardContent>
					</Card>
				</div>
			{/each}
		</div>
	</CardContent>
</Card>

<style>
	.scrollbar-hide {
		-ms-overflow-style: none;
		scrollbar-width: none;
	}
	.scrollbar-hide::-webkit-scrollbar {
		display: none;
	}
	
	.line-clamp-2 {
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
	}
</style>
