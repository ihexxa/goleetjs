/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var containsNearbyDuplicate = function(nums, k) {
		if (!nums || nums.length <=1 || k <= 0 )
			return false;

		var map = [];
		for(var i = 0; i < nums.length; i++) {
			var elem = nums[i];
			if (map[str(elem)])
				return true;
			
			map[str(elem)] = true;
			if (i - k >= 0) {
				map[str(nums[i-k])] = false;
			}
		}
		return false;
};

function str(num) {
	return num + '';
}
