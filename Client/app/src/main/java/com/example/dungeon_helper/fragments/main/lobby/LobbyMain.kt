package com.example.dungeon_helper.fragments.main.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentLobbyMainBinding

class LobbyMain : Fragment() {

    companion object {
        fun newInstance() = LobbyMain()
    }
    private lateinit var viewModel: LobbyMainViewModel


    private var _binding: FragmentLobbyMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val lobbyMainViewModel = ViewModelProvider(this)[LobbyMainViewModel::class.java]

        _binding = FragmentLobbyMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textLobby
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

        val listBtn = binding.lobbyListBtn
        val createBtn = binding.lobbyCreateBtn


        listBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_navigation_lobby_to_lobbyList)
        }

        createBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_navigation_lobby_to_lobbyCreate)
        }

    }

}