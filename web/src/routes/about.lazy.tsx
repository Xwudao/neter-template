import { createLazyFileRoute } from '@tanstack/react-router';
import classes from './about.module.scss';

const AboutPage = () => {
  return (
    <div className={classes.aboutPage}>
      {/* Hero Section */}
      <section className={classes.heroSection}>
        <div className={classes.container}>
          <h1 className={classes.heroTitle}>关于我们</h1>
          <p className={classes.heroSubtitle}>致力于为用户提供优质的产品和服务，用技术改变世界</p>
        </div>
      </section>

      {/* Company Info Section */}
      <section className={classes.companySection}>
        <div className={classes.container}>
          <div className={classes.sectionContent}>
            <div className={classes.textContent}>
              <h2 className={classes.sectionTitle}>我们的故事</h2>
              <p className={classes.description}>
                成立于2024年，我们是一家专注于创新技术解决方案的公司。
                我们相信技术的力量能够创造更美好的未来，为此我们不断探索和实践。
              </p>
              <p className={classes.description}>
                从初创团队到现在，我们始终坚持以用户为中心的理念， 致力于开发高质量、易用性强的产品。
              </p>
            </div>
            <div className={classes.statsGrid}>
              <div className={classes.statItem}>
                <div className={classes.statNumber}>100+</div>
                <div className={classes.statLabel}>满意客户</div>
              </div>
              <div className={classes.statItem}>
                <div className={classes.statNumber}>50+</div>
                <div className={classes.statLabel}>成功项目</div>
              </div>
              <div className={classes.statItem}>
                <div className={classes.statNumber}>24/7</div>
                <div className={classes.statLabel}>技术支持</div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Mission Section */}
      <section className={classes.missionSection}>
        <div className={classes.container}>
          <h2 className={classes.sectionTitle}>我们的使命</h2>
          <div className={classes.missionGrid}>
            <div className={classes.missionCard}>
              <div className={classes.missionIcon}>🚀</div>
              <h3 className={classes.missionTitle}>创新驱动</h3>
              <p className={classes.missionText}>持续探索前沿技术，为客户提供创新的解决方案</p>
            </div>
            <div className={classes.missionCard}>
              <div className={classes.missionIcon}>🎯</div>
              <h3 className={classes.missionTitle}>精益求精</h3>
              <p className={classes.missionText}>注重细节，追求完美，确保每个产品都达到最高标准</p>
            </div>
            <div className={classes.missionCard}>
              <div className={classes.missionIcon}>🤝</div>
              <h3 className={classes.missionTitle}>合作共赢</h3>
              <p className={classes.missionText}>与客户建立长期合作关系，共同成长，实现双赢</p>
            </div>
          </div>
        </div>
      </section>

      {/* Team Section */}
      <section className={classes.teamSection}>
        <div className={classes.container}>
          <h2 className={classes.sectionTitle}>核心团队</h2>
          <div className={classes.teamGrid}>
            <div className={classes.teamCard}>
              <div className={classes.teamAvatar}>👨‍💻</div>
              <h3 className={classes.teamName}>张三</h3>
              <p className={classes.teamRole}>技术总监</p>
              <p className={classes.teamDescription}>10年+开发经验，专注于全栈开发和系统架构设计</p>
            </div>
            <div className={classes.teamCard}>
              <div className={classes.teamAvatar}>👩‍💼</div>
              <h3 className={classes.teamName}>李四</h3>
              <p className={classes.teamRole}>产品经理</p>
              <p className={classes.teamDescription}>资深产品经理，擅长用户体验设计和产品规划</p>
            </div>
            <div className={classes.teamCard}>
              <div className={classes.teamAvatar}>👨‍🎨</div>
              <h3 className={classes.teamName}>王五</h3>
              <p className={classes.teamRole}>UI/UX设计师</p>
              <p className={classes.teamDescription}>专业设计师，致力于创造美观且实用的用户界面</p>
            </div>
          </div>
        </div>
      </section>

      {/* Contact Section */}
      <section className={classes.contactSection}>
        <div className={classes.container}>
          <h2 className={classes.sectionTitle}>联系我们</h2>
          <div className={classes.contactContent}>
            <div className={classes.contactInfo}>
              <div className={classes.contactItem}>
                <div className={classes.contactIcon}>📧</div>
                <div>
                  <h4>邮箱</h4>
                  <p>contact@example.com</p>
                </div>
              </div>
              <div className={classes.contactItem}>
                <div className={classes.contactIcon}>📞</div>
                <div>
                  <h4>电话</h4>
                  <p>+86 123-4567-8900</p>
                </div>
              </div>
              <div className={classes.contactItem}>
                <div className={classes.contactIcon}>📍</div>
                <div>
                  <h4>地址</h4>
                  <p>北京市朝阳区创新大街123号</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};

export const Route = createLazyFileRoute('/about')({
  component: AboutPage,
});
