class Solution {
    public int solution(int k, int[][] dungeons) {
        return getRate(k, dungeons);
    }

    private int getRate(int k, int[][] dungeons) {
        int answer = 0;
        for (int i = 0; i < dungeons.length; i++) {
            answer = Math.max(answer, tryNext(i, k, dungeons, new boolean[dungeons.length]));
        }
        return answer;
    }

    private int tryNext(int num, int k, int[][] dungeons, boolean[] cleared) {
        if (dungeons[num][0] > k) return 0;
        cleared[num] = true;
        k -= dungeons[num][1];
        int answer = 1;
        int others = 0;
        for (int i = 0; i < dungeons.length; i++) {
            if (!cleared[i]) others = Math.max(others, tryNext(i, k, dungeons, cleared));
        }
        cleared[num]=false;
        return answer+others;
    }
}