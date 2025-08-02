<script>
	import { Zap, Globe, TestTube, ExternalLink, Copy, CheckCircle, AlertCircle, Clock, DollarSign } from 'lucide-svelte';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';

	let { query, queryType } = $props();
	let network = $state('mainnet');
	let isLoading = $state(false);
	let copied = $state(false);
	let explorerData = $state(null);

	// Mock data based on query type
	const mockAddressData = {
		address: query,
		balance: 12.45,
		tokenAccounts: 8,
		transactions: 1247,
		firstSeen: '2023-08-15',
		lastActivity: '2 hours ago',
		tokens: [
			{ symbol: 'SOL', amount: 12.45, value: 1398.75 },
			{ symbol: 'USDC', amount: 250.00, value: 250.00 },
			{ symbol: 'RAY', amount: 89.32, value: 156.78 }
		]
	};

	const mockTransactionData = {
		signature: query,
		status: 'Success',
		block: 234567890,
		timestamp: '2024-01-15 14:32:18 UTC',
		fee: 0.000005,
		from: '7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU',
		to: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM',
		amount: 5.25,
		programId: '11111111111111111111111111111112'
	};

	function toggleNetwork() {
		network = network === 'mainnet' ? 'devnet' : 'mainnet';
		fetchExplorerData();
	}

	function copyToClipboard() {
		navigator.clipboard.writeText(query);
		copied = true;
		setTimeout(() => copied = false, 2000);
	}

	function fetchExplorerData() {
		isLoading = true;
		// Simulate API call to explorer.solana.com
		setTimeout(() => {
			explorerData = queryType === 'address' ? mockAddressData : mockTransactionData;
			isLoading = false;
		}, 1000);
	}

	function openInExplorer() {
		const baseUrl = network === 'mainnet' 
			? 'https://explorer.solana.com' 
			: 'https://explorer.solana.com?cluster=devnet';
		const path = queryType === 'address' ? 'address' : 'tx';
		window.open(`${baseUrl}/${path}/${query}`, '_blank');
	}

	// Initialize data
	$effect(() => {
		fetchExplorerData();
	});
</script>

<Card class="border-purple-700/30 bg-gradient-to-r from-purple-900/20 to-blue-900/20 backdrop-blur-sm shadow-sm">
	<CardContent class="p-6">
		<div class="flex flex-col sm:flex-row items-start sm:items-center gap-4">
			<div class="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-purple-500 to-blue-500 flex-shrink-0">
				<Zap class="h-5 w-5 text-white" />
			</div>
			<div class="flex-1 w-full">
				<div class="flex flex-col sm:flex-row items-start sm:items-center gap-2 mb-4">
					<h3 class="text-lg font-semibold text-white">
						Solana {queryType === 'address' ? 'Address' : 'Transaction'}
					</h3>
					<Badge variant="secondary" class="bg-purple-500/20 text-purple-400 border-purple-500/30">
						<Zap class="h-3 w-3 mr-1" />
						Blockchain
					</Badge>
					
					<!-- Network Toggle -->
					<div class="sm:ml-auto flex items-center gap-2 mt-2 sm:mt-0">
						<Button
							variant="outline"
							size="sm"
							onclick={toggleNetwork}
							class={`border-purple-600 text-purple-300 hover:text-white ${
								network === 'mainnet' 
									? 'bg-green-500/20 hover:bg-green-500/30' 
									: 'bg-orange-500/20 hover:bg-orange-500/30'
							}`}
						>
							{#if network === 'mainnet'}
								<Globe class="h-3 w-3 mr-1" />
								Mainnet
							{:else}
								<TestTube class="h-3 w-3 mr-1" />
								Devnet
							{/if}
						</Button>
						<Button
							variant="ghost"
							size="sm"
							onclick={openInExplorer}
							class="text-purple-300 hover:text-white hover:bg-purple-800/30"
						>
							<ExternalLink class="h-4 w-4" />
						</Button>
					</div>
				</div>

				<!-- Query Display -->
				<div class="bg-purple-900/30 rounded-lg p-3 border border-purple-700/30 mb-4">
					<div class="flex items-center gap-2">
						<span class="text-purple-300 text-sm font-mono break-all">{query}</span>
						<Button
							variant="ghost"
							size="sm"
							onclick={copyToClipboard}
							class="text-purple-400 hover:text-white hover:bg-purple-800/30 flex-shrink-0"
						>
							{#if copied}
								<CheckCircle class="h-3 w-3" />
							{:else}
								<Copy class="h-3 w-3" />
							{/if}
						</Button>
					</div>
				</div>

				{#if isLoading}
					<div class="flex items-center gap-3 py-8">
						<div class="h-6 w-6 animate-spin rounded-full border-2 border-purple-500/30 border-t-purple-500"></div>
						<p class="text-purple-200">Loading {network} data...</p>
					</div>
				{:else if explorerData}
					{#if queryType === 'address'}
						<!-- Address Data -->
						<div class="grid gap-4 md:grid-cols-2">
							<div class="space-y-3">
								<div class="flex items-center justify-between">
									<span class="text-purple-300 text-sm">Balance</span>
									<span class="text-white font-semibold">{explorerData.balance} SOL</span>
								</div>
								<div class="flex items-center justify-between">
									<span class="text-purple-300 text-sm">Token Accounts</span>
									<span class="text-white font-semibold">{explorerData.tokenAccounts}</span>
								</div>
								<div class="flex items-center justify-between">
									<span class="text-purple-300 text-sm">Transactions</span>
									<span class="text-white font-semibold">{explorerData.transactions.toLocaleString()}</span>
								</div>
								<div class="flex items-center justify-between">
									<span class="text-purple-300 text-sm">Last Activity</span>
									<span class="text-green-400 font-semibold">{explorerData.lastActivity}</span>
								</div>
							</div>
							
							<div class="bg-purple-900/20 rounded-lg p-3 border border-purple-700/20 mt-4 md:mt-0">
								<h4 class="text-white font-semibold mb-2 text-sm">Token Holdings</h4>
								<div class="space-y-2">
									{#each explorerData.tokens as token}
										<div class="flex items-center justify-between text-sm">
											<span class="text-purple-300">{token.symbol}</span>
											<div class="text-right">
												<div class="text-white">{token.amount}</div>
												<div class="text-purple-400 text-xs">${token.value}</div>
											</div>
										</div>
									{/each}
								</div>
							</div>
						</div>
					{:else}
						<!-- Transaction Data -->
						<div class="space-y-4">
							<div class="flex flex-wrap items-center gap-2">
								<Badge variant="secondary" class={
									explorerData.status === 'Success' 
										? 'bg-green-500/20 text-green-400 border-green-500/30'
										: 'bg-red-500/20 text-red-400 border-red-500/30'
								}>
									{#if explorerData.status === 'Success'}
										<CheckCircle class="h-3 w-3 mr-1" />
									{:else}
										<AlertCircle class="h-3 w-3 mr-1" />
									{/if}
									{explorerData.status}
								</Badge>
								<span class="text-purple-300 text-sm">Block #{explorerData.block.toLocaleString()}</span>
							</div>
							
							<div class="grid gap-3 md:grid-cols-2">
								<div class="space-y-2">
									<div class="flex items-center justify-between">
										<span class="text-purple-300 text-sm">Amount</span>
										<span class="text-white font-semibold">{explorerData.amount} SOL</span>
									</div>
									<div class="flex items-center justify-between">
										<span class="text-purple-300 text-sm">Fee</span>
										<span class="text-purple-400">{explorerData.fee} SOL</span>
									</div>
									<div class="flex items-center justify-between">
										<span class="text-purple-300 text-sm">Timestamp</span>
										<span class="text-purple-400 text-xs">{explorerData.timestamp}</span>
									</div>
								</div>
								
								<div class="space-y-2 mt-4 md:mt-0">
									<div>
										<span class="text-purple-300 text-sm block">From</span>
										<span class="text-white text-xs font-mono break-all">{explorerData.from}</span>
									</div>
									<div>
										<span class="text-purple-300 text-sm block">To</span>
										<span class="text-white text-xs font-mono break-all">{explorerData.to}</span>
									</div>
								</div>
							</div>
						</div>
					{/if}
				{/if}
			</div>
		</div>
	</CardContent>
</Card>
