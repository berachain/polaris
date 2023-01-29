// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package snapkv

func findStartIndex(strL []string, startQ string) int {
	// Modified binary search to find the very first element in >=startQ.
	if len(strL) == 0 {
		return -1
	}

	var left, right, mid int
	right = len(strL) - 1
	for left <= right {
		mid = (left + right) >> 1
		midStr := strL[mid]
		if midStr == startQ {
			// Handle condition where there might be multiple values equal to startQ.
			// We are looking for the very first value < midStL, that i+1 will be the first
			// element >= midStr.
			for i := mid - 1; i >= 0; i-- {
				if strL[i] != midStr {
					return i + 1
				}
			}
			return 0
		}
		if midStr < startQ {
			left = mid + 1
		} else { // midStrL > startQ
			right = mid - 1
		}
	}
	if left >= 0 && left < len(strL) && strL[left] >= startQ {
		return left
	}
	return -1
}

func findEndIndex(strL []string, endQ string) int {
	if len(strL) == 0 {
		return -1
	}

	// Modified binary search to find the very first element <endQ.
	var left, right, mid int
	right = len(strL) - 1
	for left <= right {
		mid = (left + right) >> 1
		midStr := strL[mid]
		if midStr == endQ {
			// Handle condition where there might be multiple values equal to startQ.
			// We are looking for the very first value < midStL, that i+1 will be the first
			// element >= midStr.
			for i := mid - 1; i >= 0; i-- {
				if strL[i] < midStr {
					return i + 1
				}
			}
			return 0
		}
		if midStr < endQ {
			left = mid + 1
		} else { // midStrL > startQ
			right = mid - 1
		}
	}

	// Binary search failed, now let's find a value less than endQ.
	for i := right; i >= 0; i-- {
		if strL[i] < endQ {
			return i
		}
	}

	return -1
}
