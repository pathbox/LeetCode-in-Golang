class Solution:
    def longestArithSeqLength(self, A: List[int]) -> int:
        dp = [{} for _ in range(len(A))]
        max_anx = 1
        for i in range(1, len(A)):
            for j in range(i):
                dp[i][A[i]-A[j]] = dp[j].get(A[i]-A[j], 1)+1
                max_anx = max(max_anx, dp[i][A[i]]-A[j])
        return max_anx
