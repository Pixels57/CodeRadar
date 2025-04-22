import './listProgrammers.css';
import { userService } from '../api.js';

function ListProgrammers({programmers, setIsDeleted}) {

  async function handleDeleteUser(userId) {
    await userService.deleteProgrammer(userId);
    setIsDeleted(true);

    alert("Programmer deleted successfully!");
  };

  return (
    <div>
      <ul>
        {
          (!programmers || programmers.length === 0)
          ? 
          <div className="container">
            <p>No programmers data available.</p> 
          </div>
          : 
          <>
            {
              programmers.map((user) => {
                // Parse skills JSON string into array
                const skills = typeof user.skills === 'string' ? JSON.parse(user.skills) : user.skills;
                
                return (
                  <div className="container">
                    <div className="imageColumn">
                      <img 
                        className='userImage' 
                        src={user.image ? `data:image/png;base64,${user.image}` : '/default-avatar.png'} 
                        alt="User Identity" 
                        onError={(e) => {
                          e.target.src = '/default-avatar.png';
                        }}
                      />
                    </div>
                    <div className="infoColumn">
                      <div className="row">
                        <h2 className="userName">{user.name}</h2>
                      </div>
                      <div className="row">
                        <p className="jobTitle">
                          <span style={{fontWeight: "bold"}}>
                            Job Title:&nbsp;
                          </span>
                          {user.jobTitle}
                        </p>
                      </div>
                      <div className="row">
                        <p className="jobTitle">
                          <span style={{fontWeight: "bold"}}>
                            Email:&nbsp;
                          </span>
                          {user.email}
                        </p>
                      </div>
                      <div className="row">
                        <div className="skills">
                          {skills.map((skill, index) => (
                            <div key={index} className="skillBox">{skill}</div>
                          ))}
                        </div>
                      </div>
                      
                      <div className="row" style={{marginTop: "20px"}}>
                        <button type="submit" className='deleteButton' onClick={() => handleDeleteUser(user.id)}>Delete User</button>
                      </div>
                    </div>
                  </div>
                );
              })
            }
          </>
        }
      </ul>
    </div>
  );
};

export default ListProgrammers;