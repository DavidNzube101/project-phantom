<script>
	import { LineChart, TrendingUp, TrendingDown, DollarSign, Clock, RefreshCw } from 'lucide-svelte';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';

	let { tokenSymbol } = $props();
	let isLoading = $state(false);

	const mockTokenData = {
		'SOL': {
			name: 'Solana',
			currentPrice: 115.23,
			change24h: 3.56, // positive change
			changePercent24h: 3.19,
			marketCap: 50.2, // in billions
			volume24h: 1.8, // in billions
			history: [
				{ price: 110, date: 'Mon' }, { price: 112, date: 'Tue' }, { price: 111, date: 'Wed' },
				{ price: 113, date: 'Thu' }, { price: 115, date: 'Fri' }, { price: 114, date: 'Sat' },
				{ price: 115.23, date: 'Sun' }
			]
		},
		'BTC': {
			name: 'Bitcoin',
			currentPrice: 68500.75,
			change24h: -1200.50, // negative change
			changePercent24h: -1.72,
			marketCap: 1.3, // in trillions
			volume24h: 25.5, // in billions
			history: [
				{ price: 69000, date: 'Mon' }, { price: 68800, date: 'Tue' }, { price: 69500, date: 'Wed' },
				{ price: 68000, date: 'Thu' }, { price: 68700, date: 'Fri' }, { price: 68200, date: 'Sat' },
				{ price: 68500.75, date: 'Sun' }
			]
		},
		'ETH': {
			name: 'Ethereum',
			currentPrice: 3820.10,
			change24h: 85.20, // positive change
			changePercent24h: 2.28,
			marketCap: 458.0, // in billions
			volume24h: 12.1, // in billions
			history: [
				{ price: 3750, date: 'Mon' }, { price: 3780, date: 'Tue' }, { price: 3760, date: 'Wed' },
				{ price: 3800, date: 'Thu' }, { price: 3810, date: 'Fri' }, { price: 3790, date: 'Sat' },
				{ price: 3820.10, date: 'Sun' }
			]
		}
	};

	let tokenData = $state(mockTokenData[tokenSymbol] || null);

	function fetchTokenData() {
		isLoading = true;
		setTimeout(() => {
			tokenData = mockTokenData[tokenSymbol] || null;
			isLoading = false;
		}, 800);
	}

	$effect(() => {
		fetchTokenData();
	});

	function getPriceChangeColor(change) {
		return change >= 0 ? 'text-green-400' : 'text-red-400';
	}

	function getPriceChangeIcon(change) {
		return change >= 0 ? TrendingUp : TrendingDown;
	}

	function getChartPoints() {
		if (!tokenData || !tokenData.history || tokenData.history.length < 2) return '';

		const prices = tokenData.history.map(d => d.price);
		const minPrice = Math.min(...prices);
		const maxPrice = Math.max(...prices);
		const priceRange = maxPrice - minPrice;

		const points = tokenData.history.map((data, i) => {
			const x = (i / (tokenData.history.length - 1)) * 100;
			const y = priceRange > 0 ? 100 - ((data.price - minPrice) / priceRange) * 100 : 50;
			return `${x},${y}`;
		}).join(' ');

		return points;
	}
</script>

<Card class="border-purple-700/30 bg-gradient-to-r from-blue-900/20 to-green-900/20 backdrop-blur-sm shadow-sm">
	<CardContent class="p-6">
		<div class="flex flex-col sm:flex-row items-start sm:items-center gap-4">
			<div class="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-blue-500 to-green-500 flex-shrink-0">
				<LineChart class="h-5 w-5 text-white" />
			</div>
			<div class="flex-1 w-full">
				<div class="flex flex-col sm:flex-row items-start sm:items-center gap-2 mb-4">
					<h3 class="text-lg font-semibold text-white">{tokenData?.name || tokenSymbol} Price Chart</h3>
					<Badge variant="secondary" class="bg-blue-500/20 text-blue-400 border-blue-500/30">
						<DollarSign class="h-3 w-3 mr-1" />
						Crypto
					</Badge>
					<Button
						variant="ghost"
						size="sm"
						onclick={fetchTokenData}
						disabled={isLoading}
						class="sm:ml-auto text-purple-300 hover:text-white hover:bg-purple-800/30"
					>
						<RefreshCw class={`h-4 w-4 ${isLoading ? 'animate-spin' : ''}`} />
					</Button>
				</div>

				{#if isLoading}
					<div class="flex items-center gap-3 py-8">
						<div class="h-6 w-6 animate-spin rounded-full border-2 border-purple-500/30 border-t-purple-500"></div>
						<p class="text-purple-200">Loading price data...</p>
					</div>
				{:else if tokenData}
					{@const SvelteComponent = getPriceChangeIcon(tokenData.change24h)}
					<div class="grid gap-4 md:grid-cols-2">
						<div class="bg-purple-900/30 rounded-lg p-4 border border-purple-700/30">
							<div class="flex items-center justify-between mb-2">
								<span class="text-purple-300 text-sm">Current Price</span>
								<span class="text-white font-semibold">{tokenSymbol}</span>
							</div>
							<div class="text-3xl font-bold text-white mb-1">${tokenData.currentPrice.toLocaleString()}</div>
							<div class="flex items-center gap-2 text-sm {getPriceChangeColor(tokenData.change24h)}">
								<SvelteComponent class="h-4 w-4" />
								<span>{tokenData.change24h.toFixed(2)} ({tokenData.changePercent24h.toFixed(2)}%)</span>
								<span class="text-purple-300 ml-auto">24h Change</span>
							</div>
						</div>
						
						<div class="bg-purple-900/30 rounded-lg p-4 border border-purple-700/30 mt-4 md:mt-0">
							<h4 class="text-white font-semibold mb-2 text-sm">Market Stats</h4>
							<div class="space-y-2">
								<div class="flex items-center justify-between text-sm">
									<span class="text-purple-300">Market Cap</span>
									<span class="text-white">${tokenData.marketCap.toLocaleString()}B</span>
								</div>
								<div class="flex items-center justify-between text-sm">
									<span class="text-purple-300">24h Volume</span>
									<span class="text-white">${tokenData.volume24h.toLocaleString()}B</span>
								</div>
								<div class="flex items-center justify-between text-sm">
									<span class="text-purple-300">Last Updated</span>
									<span class="text-purple-400 text-xs flex items-center gap-1"><Clock class="h-3 w-3" /> Just now</span>
								</div>
							</div>
						</div>
					</div>

					<div class="mt-6 bg-purple-900/30 rounded-lg p-4 border border-purple-700/30">
						<h4 class="text-white font-semibold mb-4">7-Day Price History</h4>
						<div class="relative h-40 w-full">
							<svg viewBox="0 0 100 100" preserveAspectRatio="none" class="w-full h-full">
								<polyline
									fill="none"
									stroke={tokenData.change24h >= 0 ? 'url(#gradient-green)' : 'url(#gradient-red)'}
									stroke-width="2"
									points={getChartPoints()}
								/>
								<defs>
									<linearGradient id="gradient-green" x1="0%" y1="0%" x2="100%" y2="0%">
										<stop offset="0%" stop-color="#86EFAC" />
										<stop offset="100%" stop-color="#22C55E" />
									</linearGradient>
									<linearGradient id="gradient-red" x1="0%" y1="0%" x2="100%" y2="0%">
										<stop offset="0%" stop-color="#FCA5A5" />
										<stop offset="100%" stop-color="#EF4444" />
									</linearGradient>
								</defs>
							</svg>
							<div class="absolute inset-x-0 bottom-0 flex justify-between text-xs text-purple-400 px-2 pb-1">
								{#each tokenData.history as data}
									<span>{data.date}</span>
								{/each}
							</div>
						</div>
					</div>
				{:else}
					<div class="text-purple-300 text-center py-8">No price data available for {tokenSymbol}.</div>
				{/if}
			</div>
		</div>
	</CardContent>
</Card>
