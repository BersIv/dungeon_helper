package com.example.dungeon_helper.fragments.main.lobby

import android.content.Intent
import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.lifecycle.Observer
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.MasterLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentLobbyCreateBinding
import com.example.dungeon_helper.shared.SharedViewModel

class LobbyCreate : Fragment() {

    companion object {
        fun newInstance() = LobbyCreate()
    }

    private lateinit var viewModel: LobbyCreateViewModel

    private var _binding: FragmentLobbyCreateBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val lobbyMainViewModel = ViewModelProvider(this)[LobbyCreateViewModel::class.java]

        _binding = FragmentLobbyCreateBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textLobbyCreate
        lobbyMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()

        val sharedViewModel = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        val name = binding.textFieldName.editText
        val pwd = binding.textFieldPwd.editText
        val num = binding.textFieldNumber.editText

        val backBtn = binding.backBtn
        val createBtn = binding.createBtn

        var token = ""
        sharedViewModel.token.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            println(it)
            token = it
        })




        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_lobbyCreate_to_navigation_lobby)
        }

        createBtn.setOnClickListener {

            val intent =
                Intent(activity as MainActivity, MasterLobbyActivity::class.java)

            intent.putExtra("token", token)
            intent.putExtra("name", name?.text)
            intent.putExtra("pwd", pwd?.text)
            intent.putExtra("num", num?.text)
            startActivity(intent)
        }

    }
}