// Code generated from FQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package fql

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 33, 360,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 229,
	10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 235, 10, 25, 3, 26, 3, 26, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 6, 29, 246, 10, 29, 13, 29,
	14, 29, 247, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 254, 10, 30, 12, 30, 14,
	30, 257, 11, 30, 3, 30, 3, 30, 3, 31, 6, 31, 262, 10, 31, 13, 31, 14, 31,
	263, 3, 32, 6, 32, 267, 10, 32, 13, 32, 14, 32, 268, 3, 32, 3, 32, 7, 32,
	273, 10, 32, 12, 32, 14, 32, 276, 11, 32, 3, 32, 5, 32, 279, 10, 32, 3,
	32, 3, 32, 6, 32, 283, 10, 32, 13, 32, 14, 32, 284, 3, 32, 5, 32, 288,
	10, 32, 3, 32, 6, 32, 291, 10, 32, 13, 32, 14, 32, 292, 3, 32, 5, 32, 296,
	10, 32, 3, 33, 3, 33, 5, 33, 300, 10, 33, 3, 33, 6, 33, 303, 10, 33, 13,
	33, 14, 33, 304, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48,
	3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3,
	53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58,
	3, 59, 3, 59, 3, 60, 3, 60, 2, 2, 61, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 2, 67, 2, 69,
	2, 71, 2, 73, 2, 75, 2, 77, 2, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2,
	91, 2, 93, 2, 95, 2, 97, 2, 99, 2, 101, 2, 103, 2, 105, 2, 107, 2, 109,
	2, 111, 2, 113, 2, 115, 2, 117, 2, 119, 2, 3, 2, 33, 5, 2, 11, 12, 15,
	15, 34, 34, 8, 2, 44, 44, 47, 48, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2,
	36, 36, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 50, 59, 67,
	72, 99, 104, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69,
	101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73,
	105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76,
	108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79,
	111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82,
	114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85,
	117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88,
	120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91,
	123, 123, 4, 2, 92, 92, 124, 124, 2, 347, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2,
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3,
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 3, 121, 3, 2, 2, 2, 5, 123,
	3, 2, 2, 2, 7, 125, 3, 2, 2, 2, 9, 132, 3, 2, 2, 2, 11, 139, 3, 2, 2, 2,
	13, 146, 3, 2, 2, 2, 15, 150, 3, 2, 2, 2, 17, 157, 3, 2, 2, 2, 19, 165,
	3, 2, 2, 2, 21, 172, 3, 2, 2, 2, 23, 180, 3, 2, 2, 2, 25, 185, 3, 2, 2,
	2, 27, 190, 3, 2, 2, 2, 29, 197, 3, 2, 2, 2, 31, 203, 3, 2, 2, 2, 33, 205,
	3, 2, 2, 2, 35, 207, 3, 2, 2, 2, 37, 210, 3, 2, 2, 2, 39, 212, 3, 2, 2,
	2, 41, 215, 3, 2, 2, 2, 43, 217, 3, 2, 2, 2, 45, 220, 3, 2, 2, 2, 47, 228,
	3, 2, 2, 2, 49, 234, 3, 2, 2, 2, 51, 236, 3, 2, 2, 2, 53, 238, 3, 2, 2,
	2, 55, 240, 3, 2, 2, 2, 57, 245, 3, 2, 2, 2, 59, 249, 3, 2, 2, 2, 61, 261,
	3, 2, 2, 2, 63, 295, 3, 2, 2, 2, 65, 297, 3, 2, 2, 2, 67, 306, 3, 2, 2,
	2, 69, 308, 3, 2, 2, 2, 71, 310, 3, 2, 2, 2, 73, 312, 3, 2, 2, 2, 75, 314,
	3, 2, 2, 2, 77, 316, 3, 2, 2, 2, 79, 318, 3, 2, 2, 2, 81, 320, 3, 2, 2,
	2, 83, 322, 3, 2, 2, 2, 85, 324, 3, 2, 2, 2, 87, 326, 3, 2, 2, 2, 89, 328,
	3, 2, 2, 2, 91, 330, 3, 2, 2, 2, 93, 332, 3, 2, 2, 2, 95, 334, 3, 2, 2,
	2, 97, 336, 3, 2, 2, 2, 99, 338, 3, 2, 2, 2, 101, 340, 3, 2, 2, 2, 103,
	342, 3, 2, 2, 2, 105, 344, 3, 2, 2, 2, 107, 346, 3, 2, 2, 2, 109, 348,
	3, 2, 2, 2, 111, 350, 3, 2, 2, 2, 113, 352, 3, 2, 2, 2, 115, 354, 3, 2,
	2, 2, 117, 356, 3, 2, 2, 2, 119, 358, 3, 2, 2, 2, 121, 122, 7, 42, 2, 2,
	122, 4, 3, 2, 2, 2, 123, 124, 7, 43, 2, 2, 124, 6, 3, 2, 2, 2, 125, 126,
	5, 105, 53, 2, 126, 127, 5, 77, 39, 2, 127, 128, 5, 91, 46, 2, 128, 129,
	5, 77, 39, 2, 129, 130, 5, 73, 37, 2, 130, 131, 5, 107, 54, 2, 131, 8,
	3, 2, 2, 2, 132, 133, 5, 85, 43, 2, 133, 134, 5, 95, 48, 2, 134, 135, 5,
	105, 53, 2, 135, 136, 5, 77, 39, 2, 136, 137, 5, 103, 52, 2, 137, 138,
	5, 107, 54, 2, 138, 10, 3, 2, 2, 2, 139, 140, 5, 75, 38, 2, 140, 141, 5,
	77, 39, 2, 141, 142, 5, 91, 46, 2, 142, 143, 5, 77, 39, 2, 143, 144, 5,
	107, 54, 2, 144, 145, 5, 77, 39, 2, 145, 12, 3, 2, 2, 2, 146, 147, 5, 105,
	53, 2, 147, 148, 5, 77, 39, 2, 148, 149, 5, 107, 54, 2, 149, 14, 3, 2,
	2, 2, 150, 151, 5, 109, 55, 2, 151, 152, 5, 99, 50, 2, 152, 153, 5, 75,
	38, 2, 153, 154, 5, 69, 35, 2, 154, 155, 5, 107, 54, 2, 155, 156, 5, 77,
	39, 2, 156, 16, 3, 2, 2, 2, 157, 158, 5, 69, 35, 2, 158, 159, 5, 95, 48,
	2, 159, 160, 5, 69, 35, 2, 160, 161, 5, 91, 46, 2, 161, 162, 5, 117, 59,
	2, 162, 163, 5, 119, 60, 2, 163, 164, 5, 77, 39, 2, 164, 18, 3, 2, 2, 2,
	165, 166, 5, 77, 39, 2, 166, 167, 5, 115, 58, 2, 167, 168, 5, 99, 50, 2,
	168, 169, 5, 97, 49, 2, 169, 170, 5, 103, 52, 2, 170, 171, 5, 107, 54,
	2, 171, 20, 3, 2, 2, 2, 172, 173, 5, 77, 39, 2, 173, 174, 5, 115, 58, 2,
	174, 175, 5, 77, 39, 2, 175, 176, 5, 73, 37, 2, 176, 177, 5, 109, 55, 2,
	177, 178, 5, 107, 54, 2, 178, 179, 5, 77, 39, 2, 179, 22, 3, 2, 2, 2, 180,
	181, 5, 85, 43, 2, 181, 182, 5, 95, 48, 2, 182, 183, 5, 107, 54, 2, 183,
	184, 5, 97, 49, 2, 184, 24, 3, 2, 2, 2, 185, 186, 5, 79, 40, 2, 186, 187,
	5, 103, 52, 2, 187, 188, 5, 97, 49, 2, 188, 189, 5, 93, 47, 2, 189, 26,
	3, 2, 2, 2, 190, 191, 5, 111, 56, 2, 191, 192, 5, 69, 35, 2, 192, 193,
	5, 91, 46, 2, 193, 194, 5, 109, 55, 2, 194, 195, 5, 77, 39, 2, 195, 196,
	5, 105, 53, 2, 196, 28, 3, 2, 2, 2, 197, 198, 5, 113, 57, 2, 198, 199,
	5, 83, 42, 2, 199, 200, 5, 77, 39, 2, 200, 201, 5, 103, 52, 2, 201, 202,
	5, 77, 39, 2, 202, 30, 3, 2, 2, 2, 203, 204, 7, 44, 2, 2, 204, 32, 3, 2,
	2, 2, 205, 206, 7, 63, 2, 2, 206, 34, 3, 2, 2, 2, 207, 208, 7, 63, 2, 2,
	208, 209, 7, 63, 2, 2, 209, 36, 3, 2, 2, 2, 210, 211, 7, 62, 2, 2, 211,
	38, 3, 2, 2, 2, 212, 213, 7, 62, 2, 2, 213, 214, 7, 63, 2, 2, 214, 40,
	3, 2, 2, 2, 215, 216, 7, 64, 2, 2, 216, 42, 3, 2, 2, 2, 217, 218, 7, 64,
	2, 2, 218, 219, 7, 63, 2, 2, 219, 44, 3, 2, 2, 2, 220, 221, 7, 35, 2, 2,
	221, 222, 7, 63, 2, 2, 222, 46, 3, 2, 2, 2, 223, 224, 5, 69, 35, 2, 224,
	225, 5, 95, 48, 2, 225, 226, 5, 75, 38, 2, 226, 229, 3, 2, 2, 2, 227, 229,
	7, 40, 2, 2, 228, 223, 3, 2, 2, 2, 228, 227, 3, 2, 2, 2, 229, 48, 3, 2,
	2, 2, 230, 231, 5, 97, 49, 2, 231, 232, 5, 103, 52, 2, 232, 235, 3, 2,
	2, 2, 233, 235, 7, 126, 2, 2, 234, 230, 3, 2, 2, 2, 234, 233, 3, 2, 2,
	2, 235, 50, 3, 2, 2, 2, 236, 237, 7, 46, 2, 2, 237, 52, 3, 2, 2, 2, 238,
	239, 7, 61, 2, 2, 239, 54, 3, 2, 2, 2, 240, 241, 9, 2, 2, 2, 241, 242,
	3, 2, 2, 2, 242, 243, 8, 28, 2, 2, 243, 56, 3, 2, 2, 2, 244, 246, 9, 3,
	2, 2, 245, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2,
	247, 248, 3, 2, 2, 2, 248, 58, 3, 2, 2, 2, 249, 255, 7, 36, 2, 2, 250,
	251, 7, 36, 2, 2, 251, 254, 7, 36, 2, 2, 252, 254, 10, 4, 2, 2, 253, 250,
	3, 2, 2, 2, 253, 252, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253, 3, 2,
	2, 2, 255, 256, 3, 2, 2, 2, 256, 258, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2,
	258, 259, 7, 36, 2, 2, 259, 60, 3, 2, 2, 2, 260, 262, 4, 50, 59, 2, 261,
	260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 263, 264,
	3, 2, 2, 2, 264, 62, 3, 2, 2, 2, 265, 267, 4, 50, 59, 2, 266, 265, 3, 2,
	2, 2, 267, 268, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2,
	269, 270, 3, 2, 2, 2, 270, 274, 7, 48, 2, 2, 271, 273, 4, 50, 59, 2, 272,
	271, 3, 2, 2, 2, 273, 276, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 274, 275,
	3, 2, 2, 2, 275, 278, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 277, 279, 5, 65,
	33, 2, 278, 277, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 296, 3, 2, 2, 2,
	280, 282, 7, 48, 2, 2, 281, 283, 4, 50, 59, 2, 282, 281, 3, 2, 2, 2, 283,
	284, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 287,
	3, 2, 2, 2, 286, 288, 5, 65, 33, 2, 287, 286, 3, 2, 2, 2, 287, 288, 3,
	2, 2, 2, 288, 296, 3, 2, 2, 2, 289, 291, 4, 50, 59, 2, 290, 289, 3, 2,
	2, 2, 291, 292, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2,
	293, 294, 3, 2, 2, 2, 294, 296, 5, 65, 33, 2, 295, 266, 3, 2, 2, 2, 295,
	280, 3, 2, 2, 2, 295, 290, 3, 2, 2, 2, 296, 64, 3, 2, 2, 2, 297, 299, 9,
	5, 2, 2, 298, 300, 9, 6, 2, 2, 299, 298, 3, 2, 2, 2, 299, 300, 3, 2, 2,
	2, 300, 302, 3, 2, 2, 2, 301, 303, 4, 50, 59, 2, 302, 301, 3, 2, 2, 2,
	303, 304, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305,
	66, 3, 2, 2, 2, 306, 307, 9, 7, 2, 2, 307, 68, 3, 2, 2, 2, 308, 309, 9,
	8, 2, 2, 309, 70, 3, 2, 2, 2, 310, 311, 9, 9, 2, 2, 311, 72, 3, 2, 2, 2,
	312, 313, 9, 10, 2, 2, 313, 74, 3, 2, 2, 2, 314, 315, 9, 11, 2, 2, 315,
	76, 3, 2, 2, 2, 316, 317, 9, 5, 2, 2, 317, 78, 3, 2, 2, 2, 318, 319, 9,
	12, 2, 2, 319, 80, 3, 2, 2, 2, 320, 321, 9, 13, 2, 2, 321, 82, 3, 2, 2,
	2, 322, 323, 9, 14, 2, 2, 323, 84, 3, 2, 2, 2, 324, 325, 9, 15, 2, 2, 325,
	86, 3, 2, 2, 2, 326, 327, 9, 16, 2, 2, 327, 88, 3, 2, 2, 2, 328, 329, 9,
	17, 2, 2, 329, 90, 3, 2, 2, 2, 330, 331, 9, 18, 2, 2, 331, 92, 3, 2, 2,
	2, 332, 333, 9, 19, 2, 2, 333, 94, 3, 2, 2, 2, 334, 335, 9, 20, 2, 2, 335,
	96, 3, 2, 2, 2, 336, 337, 9, 21, 2, 2, 337, 98, 3, 2, 2, 2, 338, 339, 9,
	22, 2, 2, 339, 100, 3, 2, 2, 2, 340, 341, 9, 23, 2, 2, 341, 102, 3, 2,
	2, 2, 342, 343, 9, 24, 2, 2, 343, 104, 3, 2, 2, 2, 344, 345, 9, 25, 2,
	2, 345, 106, 3, 2, 2, 2, 346, 347, 9, 26, 2, 2, 347, 108, 3, 2, 2, 2, 348,
	349, 9, 27, 2, 2, 349, 110, 3, 2, 2, 2, 350, 351, 9, 28, 2, 2, 351, 112,
	3, 2, 2, 2, 352, 353, 9, 29, 2, 2, 353, 114, 3, 2, 2, 2, 354, 355, 9, 30,
	2, 2, 355, 116, 3, 2, 2, 2, 356, 357, 9, 31, 2, 2, 357, 118, 3, 2, 2, 2,
	358, 359, 9, 32, 2, 2, 359, 120, 3, 2, 2, 2, 18, 2, 228, 234, 247, 253,
	255, 263, 268, 274, 278, 284, 287, 292, 295, 299, 304, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "", "", "", "", "", "", "", "", "", "", "", "", "'*'",
	"'='", "'=='", "'<'", "'<='", "'>'", "'>='", "'!='", "", "", "','", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "", "SELECT", "INSERT", "DELETE", "SET", "UPDATE", "ANALYZE", "EXPORT",
	"EXECUTE", "INTO", "FROM", "VALUES", "WHERE", "ASTERISK", "SINGLE_EQ",
	"DOUBLE_EQ", "OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR", "COMMA",
	"SEMICOLON", "WS", "IDENTIFIER", "STRING", "NUMBER", "REAL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "SELECT", "INSERT", "DELETE", "SET", "UPDATE", "ANALYZE",
	"EXPORT", "EXECUTE", "INTO", "FROM", "VALUES", "WHERE", "ASTERISK", "SINGLE_EQ",
	"DOUBLE_EQ", "OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR", "COMMA",
	"SEMICOLON", "WS", "IDENTIFIER", "STRING", "NUMBER", "REAL", "EXPONENT",
	"HEX_DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

type FQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewFQLLexer(input antlr.CharStream) *FQLLexer {

	l := new(FQLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "FQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FQLLexer tokens.
const (
	FQLLexerT__0       = 1
	FQLLexerT__1       = 2
	FQLLexerSELECT     = 3
	FQLLexerINSERT     = 4
	FQLLexerDELETE     = 5
	FQLLexerSET        = 6
	FQLLexerUPDATE     = 7
	FQLLexerANALYZE    = 8
	FQLLexerEXPORT     = 9
	FQLLexerEXECUTE    = 10
	FQLLexerINTO       = 11
	FQLLexerFROM       = 12
	FQLLexerVALUES     = 13
	FQLLexerWHERE      = 14
	FQLLexerASTERISK   = 15
	FQLLexerSINGLE_EQ  = 16
	FQLLexerDOUBLE_EQ  = 17
	FQLLexerOP_LT      = 18
	FQLLexerLE         = 19
	FQLLexerGT         = 20
	FQLLexerGE         = 21
	FQLLexerNOTEQ      = 22
	FQLLexerAND        = 23
	FQLLexerOR         = 24
	FQLLexerCOMMA      = 25
	FQLLexerSEMICOLON  = 26
	FQLLexerWS         = 27
	FQLLexerIDENTIFIER = 28
	FQLLexerSTRING     = 29
	FQLLexerNUMBER     = 30
	FQLLexerREAL       = 31
)
