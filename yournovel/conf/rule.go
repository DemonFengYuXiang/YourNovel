package conf

type Config struct {
	Engines      []string
	Rule         map[string]map[string]interface{}
	IgnoreDomain map[string]int
}

var RuleConfig *Config

func InitConfig() {

	RuleConfig = &Config{}

	RuleConfig.Engines = []string{"baidu"}

	// link_prefix -1 代表取chapter_url 1 代表直接取URL 0 代表使用域名加拼接
	RuleConfig.Rule = map[string]map[string]interface{}{
		"www.biquge.info": {"link_prefix": "-1", "chapter_selector": "#list", "content_selector": "#content",},
		"www.biqugetv.com": {"link_prefix":"0", "chapter_selector": "#list", "content_selector": "#content",},
		"www.paoshu8.com": {"link_prefix":"0", "chapter_selector": "#list", "content_selector": "#content",},
		"www.xbiquge.la": {"link_prefix":"0", "chapter_selector": "#list", "content_selector": "#content",},
		"www.x23us.com": {"link_prefix":"-1", "chapter_selector": "#at", "content_selector": "#contents",},
		"www.81zw.us": {"link_prefix":"0", "chapter_selector": "#list", "content_selector": "#content",},
		"www.022003.com": {"link_prefix":"0", "chapter_selector": "#list", "content_selector": "#content",},
		"www.biquguan.com": {"link_prefix":"0", "chapter_selector": "#list", "content_selector": "#content",},
		"www.xqqxs.com": {"link_prefix":"-1", "chapter_selector": "#list", "content_selector": "#booktext",},
		"www.bbiquge.com": {"link_prefix":"-1", "chapter_selector": "#list", "content_selector": "#content",},
		"www.xshengyan.com": {"link_prefix":"0", "chapter_selector": "#list", "content_selector": "#content",},
		"www.biquge.tv": {"link_prefix":"-1", "chapter_selector": "#list", "content_selector": "#content",},
	}

	RuleConfig.IgnoreDomain = map[string]int {
		"www.17k.com":1, "mm.17k.com":1, "www.xs8.cn":1, "www.zongheng.com":1, "yunqi.qq.com":1, "chuangshi.qq.com":1,
		"book.qidian.com":1, "www.soduso.com":1, "pages.book.qq.com":1, "book.km.com":1, "www.lread.net":1,
		"www.0dsw.com":1, "www.5200xsb.com":1, "www.80txt.com":1, "www.sodu.tw":1, "www.shuquge.com":1,
		"www.shenmanhua.com":1, "xiaoshuo.sogou.com":1, "www.999wx.com":1, "zetianji8.com":1, "www.bookso.net":1,
		"m.23us.com":1, "www.qbxsw.com":1, "www.zhuzhudao.com":1, "www.shengyan.org":1, "www.360doc.com":1,
		"www.ishuo.cn":1, "read.qidian.com":1, "www.yunlaige.com":1, "www.qidian.com":1, "www.sodu888.com":1,
		"www.siluke.cc":1, "read.10086.cn":1, "www.pbtxt.com":1, "c4txt.com":1, "www.bokon.net":1, "www.sikushu.net":1,
		"www.is028.cn":1, "www.tadu.com":1, "www.kudu8.com":1, "www.bmwen.com":1, "www.5858xs.com":1, "www.yiwan.com":1,
		"www.x81zw.com":1, "www.123du.cc":1, "www.chashu.cc":1, "20xs.com":1, "www.haxwx.net":1, "www.dushiwenxue.com":1,
		"www.yxdown.com":1, "www.jingcaiyuedu.com":1, "www.zhetian.org":1, "www.xiaoshuo02.com":1, "www.xiaoshuo77.com":1,
		"www.868xh.com":1, "dp.changyou.com":1, "www.iyouman.com":1, "www.qq717.com":1, "www.yznn.com":1, "www.69w.cc":1,
		"www.doupocangqiong1.com":1, "www.manhuatai.com":1, "www.5wxs.com":1, "www.ggshuji.com":1, "www.msxf.net":1,
		"www.mianhuatang.la":1, "www.boluoxs.com":1, "www.lbiquge.top":1, "www.69shu.com":1, "www.qingkan520.com":1,
		"book.douban.com":1, "movie.douban.com":1, "www.txshuku.com":1, "lz.book.sohu.com":1, "www.3gsc.com.cn":1,
		"www.txtshu365.com":1, "www.517yuedu.com":1, "www.baike.com":1, "read.jd.com":1, "www.zhihu.com":1, "wshuyi.com":1,
		"www.19lou.tw":1, "www.chenwangbook.com":1, "www.aqtxt.com":1, "book.114la.com":1, "www.niepo.net":1,
		"me.qidian.com":1, "www.gengd.com":1, "www.77l.com":1, "www.geilwx.com":1, "www.97xiao.com":1, "www.anqu.com":1,
		"www.wuxiaxs.com":1, "yuedu.163.com":1, "b.faloo.com":1, "bbs.qidian.com":1, "jingji.qidian.com":1, "www.sodu.cc":1,
		"forum.qdmm.com":1, "www.qdmm.com":1, "game.91.com":1, "www.11773.com":1, "mt.sohu.com":1, "book.dajianet.com":1,
		"haokan.17k.com":1, "www.qmdsj.com":1, "www.jjwxc.net":1, "ishare.iask.sina.com.cn":1, "www.cmread.com":1,
		"www.52ranwen.net":1, "www.dingdianzw.com":1, "www.topber.com":1, "www.391k.com":1, "www.qqxzb.com":1,
		"www.zojpw.com":1, "www.pp8.com":1, "www.bxwx.org":1, "www.hrsxb.com":1, "www.497.com":1, "www.d8qu.com":1,
		"www.duwanjuan.com":1, "www.05935.com":1, "book.zongheng.com":1, "www.55x.cn":1, "www.freexs.cn":1,
		"xiaoshuo.360.cn":1, "www.3kw.cc":1, "www.gzbpi.com":1, "book.sina.com.cn":1, "www.vodtw.com":1, "wenda.so.com":1,
		"product.dangdang.com":1, "www.chuiyao.com":1, "novel.slieny.com":1, "www.bilibili.com":1, "donghua.dmzj.com":1,
		"www.yaojingweiba.com":1, "www.qb5200.com":1, "www.520tingshu.com":1, "www.567zw.com":1, "www.zjrxz.com":1,
		"v.qq.com":1, "blog.sina.com.cn":1, "www.hackhome.com":1, "news.fznews.com.cn":1, "www.jingyu.com":1,
		"news.so.com":1, "www.sodu3.com":1, "vipreader.qidian.com":1, "www.mozhua9.com":1, "www.iqiyi.com":1,
		"xs.sogou.com":1, "www.xiashuwu.com":1,
	}
}
