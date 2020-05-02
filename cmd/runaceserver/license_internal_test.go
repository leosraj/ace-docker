/*
© Copyright IBM Corporation 2018

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE_2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"os"
	"testing"
)

var licenseTests = []struct {
	in  string
	out string
}{
	{"en_US.UTF_8", "English.txt"},
	{"en_US.ISO-8859-15", "English.txt"},
	{"es_GB", "Spanish.txt"},
	{"el_ES.UTF_8", "Greek.txt"},
	// Cover a wide variety of valid values
	{"af", "English.txt"},
	{"af_ZA", "English.txt"},
	{"ar", "English.txt"},
	{"ar_AE", "English.txt"},
	{"ar_BH", "English.txt"},
	{"ar_DZ", "English.txt"},
	{"ar_EG", "English.txt"},
	{"ar_IQ", "English.txt"},
	{"ar_JO", "English.txt"},
	{"ar_KW", "English.txt"},
	{"ar_LB", "English.txt"},
	{"ar_LY", "English.txt"},
	{"ar_MA", "English.txt"},
	{"ar_OM", "English.txt"},
	{"ar_QA", "English.txt"},
	{"ar_SA", "English.txt"},
	{"ar_SY", "English.txt"},
	{"ar_TN", "English.txt"},
	{"ar_YE", "English.txt"},
	{"az", "English.txt"},
	{"az_AZ", "English.txt"},
	{"az_AZ", "English.txt"},
	{"be", "English.txt"},
	{"be_BY", "English.txt"},
	{"bg", "English.txt"},
	{"bg_BG", "English.txt"},
	{"bs_BA", "English.txt"},
	{"ca", "English.txt"},
	{"ca_ES", "English.txt"},
	{"cs", "Czech.txt"},
	{"cs_CZ", "Czech.txt"},
	{"csb_PL", "English.txt"},
	{"cy", "English.txt"},
	{"cy_GB", "English.txt"},
	{"da", "English.txt"},
	{"da_DK", "English.txt"},
	{"de", "German.txt"},
	{"de_AT", "German.txt"},
	{"de_CH", "German.txt"},
	{"de_DE", "German.txt"},
	{"de_LI", "German.txt"},
	{"de_LU", "German.txt"},
	{"dv", "English.txt"},
	{"dv_MV", "English.txt"},
	{"el", "Greek.txt"},
	{"el_GR", "Greek.txt"},
	{"en", "English.txt"},
	{"en_AU", "English.txt"},
	{"en_BZ", "English.txt"},
	{"en_CA", "English.txt"},
	{"en_CB", "English.txt"},
	{"en_GB", "English.txt"},
	{"en_IE", "English.txt"},
	{"en_JM", "English.txt"},
	{"en_NZ", "English.txt"},
	{"en_PH", "English.txt"},
	{"en_TT", "English.txt"},
	{"en_US", "English.txt"},
	{"en_ZA", "English.txt"},
	{"en_ZW", "English.txt"},
	{"eo", "English.txt"},
	{"es", "Spanish.txt"},
	{"es_AR", "Spanish.txt"},
	{"es_BO", "Spanish.txt"},
	{"es_CL", "Spanish.txt"},
	{"es_CO", "Spanish.txt"},
	{"es_CR", "Spanish.txt"},
	{"es_DO", "Spanish.txt"},
	{"es_EC", "Spanish.txt"},
	{"es_ES", "Spanish.txt"},
	{"es_ES", "Spanish.txt"},
	{"es_GT", "Spanish.txt"},
	{"es_HN", "Spanish.txt"},
	{"es_MX", "Spanish.txt"},
	{"es_NI", "Spanish.txt"},
	{"es_PA", "Spanish.txt"},
	{"es_PE", "Spanish.txt"},
	{"es_PR", "Spanish.txt"},
	{"es_PY", "Spanish.txt"},
	{"es_SV", "Spanish.txt"},
	{"es_UY", "Spanish.txt"},
	{"es_VE", "Spanish.txt"},
	{"et", "English.txt"},
	{"et_EE", "English.txt"},
	{"eu", "English.txt"},
	{"eu_ES", "English.txt"},
	{"fa", "English.txt"},
	{"fa_IR", "English.txt"},
	{"fi", "English.txt"},
	{"fi_FI", "English.txt"},
	{"fo", "English.txt"},
	{"fo_FO", "English.txt"},
	{"fr", "French.txt"},
	{"fr_BE", "French.txt"},
	{"fr_CA", "French.txt"},
	{"fr_CH", "French.txt"},
	{"fr_FR", "French.txt"},
	{"fr_LU", "French.txt"},
	{"fr_MC", "French.txt"},
	{"gl", "English.txt"},
	{"gl_ES", "English.txt"},
	{"gu", "English.txt"},
	{"gu_IN", "English.txt"},
	{"he", "English.txt"},
	{"he_IL", "English.txt"},
	{"hi", "English.txt"},
	{"hi_IN", "English.txt"},
	{"hr", "English.txt"},
	{"hr_BA", "English.txt"},
	{"hr_HR", "English.txt"},
	{"hu", "English.txt"},
	{"hu_HU", "English.txt"},
	{"hy", "English.txt"},
	{"hy_AM", "English.txt"},
	{"id", "Indonesian.txt"},
	{"id_ID", "Indonesian.txt"},
	{"is", "English.txt"},
	{"is_IS", "English.txt"},
	{"it", "Italian.txt"},
	{"it_CH", "Italian.txt"},
	{"it_IT", "Italian.txt"},
	{"ja", "Japanese.txt"},
	{"ja_JP", "Japanese.txt"},
	{"ka", "English.txt"},
	{"ka_GE", "English.txt"},
	{"kk", "English.txt"},
	{"kk_KZ", "English.txt"},
	{"kn", "English.txt"},
	{"kn_IN", "English.txt"},
	{"ko", "Korean.txt"},
	{"ko_KR", "Korean.txt"},
	{"kok", "English.txt"},
	{"kok_IN", "English.txt"},
	{"ky", "English.txt"},
	{"ky_KG", "English.txt"},
	{"lt", "Lithuanian.txt"},
	{"lt_LT", "Lithuanian.txt"},
	{"lv", "English.txt"},
	{"lv_LV", "English.txt"},
	{"mi", "English.txt"},
	{"mi_NZ", "English.txt"},
	{"mk", "English.txt"},
	{"mk_MK", "English.txt"},
	{"mn", "English.txt"},
	{"mn_MN", "English.txt"},
	{"mr", "English.txt"},
	{"mr_IN", "English.txt"},
	{"ms", "English.txt"},
	{"ms_BN", "English.txt"},
	{"ms_MY", "English.txt"},
	{"mt", "English.txt"},
	{"mt_MT", "English.txt"},
	{"nb", "English.txt"},
	{"nb_NO", "English.txt"},
	{"nl", "English.txt"},
	{"nl_BE", "English.txt"},
	{"nl_NL", "English.txt"},
	{"nn_NO", "English.txt"},
	{"ns", "English.txt"},
	{"ns_ZA", "English.txt"},
	{"pa", "English.txt"},
	{"pa_IN", "English.txt"},
	{"pl", "Polish.txt"},
	{"pl_PL", "Polish.txt"},
	{"ps", "English.txt"},
	{"ps_AR", "English.txt"},
	{"pt", "Portugese.txt"},
	{"pt_BR", "Portugese.txt"},
	{"pt_PT", "Portugese.txt"},
	{"qu", "English.txt"},
	{"qu_BO", "English.txt"},
	{"qu_EC", "English.txt"},
	{"qu_PE", "English.txt"},
	{"ro", "English.txt"},
	{"ro_RO", "English.txt"},
	{"ru", "Russian.txt"},
	{"ru_RU", "Russian.txt"},
	{"sa", "English.txt"},
	{"sa_IN", "English.txt"},
	{"se", "English.txt"},
	{"se_FI", "English.txt"},
	{"se_FI", "English.txt"},
	{"se_FI", "English.txt"},
	{"se_NO", "English.txt"},
	{"se_NO", "English.txt"},
	{"se_NO", "English.txt"},
	{"se_SE", "English.txt"},
	{"se_SE", "English.txt"},
	{"se_SE", "English.txt"},
	{"sk", "English.txt"},
	{"sk_SK", "English.txt"},
	{"sl", "Slovenian.txt"},
	{"sl_SI", "Slovenian.txt"},
	{"sq", "English.txt"},
	{"sq_AL", "English.txt"},
	{"sr_BA", "English.txt"},
	{"sr_BA", "English.txt"},
	{"sr_SP", "English.txt"},
	{"sr_SP", "English.txt"},
	{"sv", "English.txt"},
	{"sv_FI", "English.txt"},
	{"sv_SE", "English.txt"},
	{"sw", "English.txt"},
	{"sw_KE", "English.txt"},
	{"syr", "English.txt"},
	{"syr_SY", "English.txt"},
	{"ta", "English.txt"},
	{"ta_IN", "English.txt"},
	{"te", "English.txt"},
	{"te_IN", "English.txt"},
	{"th", "English.txt"},
	{"th_TH", "English.txt"},
	{"tl", "English.txt"},
	{"tl_PH", "English.txt"},
	{"tn", "English.txt"},
	{"tn_ZA", "English.txt"},
	{"tr", "Turkish.txt"},
	{"tr_TR", "Turkish.txt"},
	{"tt", "English.txt"},
	{"tt_RU", "English.txt"},
	{"ts", "English.txt"},
	{"uk", "English.txt"},
	{"uk_UA", "English.txt"},
	{"ur", "English.txt"},
	{"ur_PK", "English.txt"},
	{"uz", "English.txt"},
	{"uz_UZ", "English.txt"},
	{"uz_UZ", "English.txt"},
	{"vi", "English.txt"},
	{"vi_VN", "English.txt"},
	{"xh", "English.txt"},
	{"xh_ZA", "English.txt"},
	{"zh", "Chinese.txt"},
	{"zh_CN", "Chinese.txt"},
	{"zh_HK", "Chinese.txt"},
	{"zh_MO", "Chinese.txt"},
	{"zh_SG", "Chinese.txt"},
	{"zh_TW", "Chinese_TW.txt"},
	{"zu", "English.txt"},
	{"zu_ZA", "English.txt"},
}

func TestResolveLicenseFile(t *testing.T) {
	for _, table := range licenseTests {
		os.Setenv("LANG", table.in)
		f := resolveLicenseFile()
		if f != table.out {
			t.Errorf("resolveLicenseFile() with LANG=%v - expected %v, got %v", table.in, table.out, f)
		}
	}
}
